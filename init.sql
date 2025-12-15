CREATE TABLE IF NOT EXISTS notes (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_notes_title_gin 
ON notes USING GIN (to_tsvector('simple', title));

CREATE INDEX IF NOT EXISTS idx_notes_created_id 
ON notes (created_at DESC, id DESC);

CREATE EXTENSION IF NOT EXISTS pg_stat_statements;

INSERT INTO notes (title, content) VALUES
('Первая заметка', 'Содержание первой заметки'),
('Вторая заметка', 'Содержание второй заметки'),
('Третья заметка', 'Содержание третьей заметки'),
('Четвертая заметка', 'Содержание четвертой заметки'),
('Пятая заметка', 'Содержание пятой заметки'),
('Шестая заметка', 'Содержание шестой заметки'),
('Седьмая заметка', 'Содержание седьмой заметки'),
('Восьмая заметка', 'Содержание восьмой заметки'),
('Девятая заметка', 'Содержание девятой заметки'),
('Десятая заметка', 'Содержание десятой заметки')
ON CONFLICT DO NOTHING;