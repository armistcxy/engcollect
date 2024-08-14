CREATE TABLE IF NOT EXISTS words(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    word TEXT NOT NULL,
    level TEXT,
    definition TEXT NOT NULL,
    example_usage TEXT,
    optinal_note TEXT,
    created_at DATE DEFAULT(date('now')),
    updated_at DATE DEFAULT(date('now'))
);


-- create index on word in "words" table if index on word has not created yet
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM sqlite_master WHERE type='index' AND name = 'idx_word') THEN
        CREATE INDEX idx_word ON words(word); 
    ENDIF;
END $$; 

