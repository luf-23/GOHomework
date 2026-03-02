package main

/*
1.文件读取：从 json 目录内挑选以下两个 json 文件。下载到本地仓库中，直接 git 提交。
（图示路径：english-vocabulary / json /，选中文件为 3-CET4-顺序.json 和 4-CET6-顺序.json）
2.数据解析：对读取的 JSON 文件进行解析，提取其中的单词数据。
3.字段拆分：将解析后的单词数据，按照不同字段进行拆分。
明确每个字段的含义和数据类型。例如，常见字段可能有 “单词 (word)”、“短语 (phrases)”、“词性 (translations)” 等。
可根据自己理解灵活拆分，但是不能有信息丢失和冗余。（数据库字段可存储 json 信息）
4.数据库写入：将拆分后的数据写入 SQLite 数据库（两个文件写入一个数据库）。
需要正确创建数据库表结构，保证数据能够准确无误地插入到相应的表和字段中。sqlite 数据库查看可以用 sqliteviewer 插件或者 dbeaver（推荐）
*/

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

// WordEntry 对应 JSON 中的一个单词条目
type WordEntry struct {
	Word         string        `json:"word"`
	Translations []Translation `json:"translations"`
	Phrases      []Phrase      `json:"phrases"`
}

// Translation 单词的一个词义（含词性）
type Translation struct {
	Translation string `json:"translation"`
	Type        string `json:"type"` // 词性：adj/n/v/vt/vi/adv 等
}

// Phrase 单词的一个短语
type Phrase struct {
	Phrase      string `json:"phrase"`
	Translation string `json:"translation"`
}

// parseFile 读取并解析 JSON 文件，返回单词列表
func parseFile(path string) ([]WordEntry, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取文件失败 %s: %w", path, err)
	}
	var entries []WordEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, fmt.Errorf("解析 JSON 失败 %s: %w", path, err)
	}
	return entries, nil
}

// initDB 初始化数据库，创建三张表
func initDB(db *sql.DB) error {
	schema := `
-- 单词主表：存储单词本身及来源词库
CREATE TABLE IF NOT EXISTS words (
    id     INTEGER PRIMARY KEY AUTOINCREMENT,
    word   TEXT    NOT NULL,
    source TEXT    NOT NULL  -- CET4 或 CET6
);

-- 词义表：一个单词可有多个词义/词性
CREATE TABLE IF NOT EXISTS translations (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    word_id     INTEGER NOT NULL REFERENCES words(id),
    translation TEXT    NOT NULL,  -- 中文释义
    type        TEXT               -- 词性（adj/n/v/vt/vi/adv 等），可能为空
);

-- 短语表：一个单词可有多个相关短语
CREATE TABLE IF NOT EXISTS phrases (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    word_id     INTEGER NOT NULL REFERENCES words(id),
    phrase      TEXT    NOT NULL,  -- 英文短语
    translation TEXT    NOT NULL   -- 短语中文释义
);
`
	_, err := db.Exec(schema)
	return err
}

// insertEntries 将解析好的单词条目批量写入数据库
func insertEntries(db *sql.DB, entries []WordEntry, source string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmtWord, err := tx.Prepare(`INSERT INTO words (word, source) VALUES (?, ?)`)
	if err != nil {
		return err
	}
	defer stmtWord.Close()

	stmtTrans, err := tx.Prepare(`INSERT INTO translations (word_id, translation, type) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmtTrans.Close()

	stmtPhrase, err := tx.Prepare(`INSERT INTO phrases (word_id, phrase, translation) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmtPhrase.Close()

	for _, entry := range entries {
		res, err := stmtWord.Exec(entry.Word, source)
		if err != nil {
			return fmt.Errorf("插入单词 %q 失败: %w", entry.Word, err)
		}
		wordID, _ := res.LastInsertId()

		for _, t := range entry.Translations {
			if _, err := stmtTrans.Exec(wordID, t.Translation, t.Type); err != nil {
				return fmt.Errorf("插入词义失败 (word=%q): %w", entry.Word, err)
			}
		}

		for _, p := range entry.Phrases {
			if _, err := stmtPhrase.Exec(wordID, p.Phrase, p.Translation); err != nil {
				return fmt.Errorf("插入短语失败 (word=%q, phrase=%q): %w", entry.Word, p.Phrase, err)
			}
		}
	}

	return tx.Commit()
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取工作目录失败: %v", err)
	}

	sources := []struct {
		path   string
		source string
	}{
		{filepath.Join(wd, "resource", "CET4.json"), "CET4"},
		{filepath.Join(wd, "resource", "CET6.json"), "CET6"},
	}

	dbPath := filepath.Join(wd, "resource", "vocabulary.db")

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("打开数据库失败: %v", err)
	}
	defer db.Close()

	if err := initDB(db); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	for _, s := range sources {
		entries, err := parseFile(s.path)
		if err != nil {
			log.Fatalf("解析失败: %v", err)
		}
		if err := insertEntries(db, entries, s.source); err != nil {
			log.Fatalf("写入 %s 失败: %v", s.source, err)
		}
		fmt.Printf("✓ %s: 共写入 %d 个单词\n", s.source, len(entries))
	}

	var total int
	if err := db.QueryRow(`SELECT COUNT(*) FROM words`).Scan(&total); err != nil {
		log.Fatalf("统计失败: %v", err)
	}
	fmt.Printf("✓ 数据库共 %d 个单词\n", total)
}
