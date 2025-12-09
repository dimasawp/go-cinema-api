-- ===============================
-- DROP TABLE (anak → induk)
-- ===============================
DROP TABLE IF EXISTS refunds CASCADE;
DROP TABLE IF EXISTS tickets CASCADE;
DROP TABLE IF EXISTS transactions CASCADE;
DROP TABLE IF EXISTS seat_locks CASCADE;
DROP TABLE IF EXISTS showtimes CASCADE;
DROP TABLE IF EXISTS movies CASCADE;
DROP TABLE IF EXISTS seats CASCADE;
DROP TABLE IF EXISTS auditoriums CASCADE;
DROP TABLE IF EXISTS cinemas CASCADE;
DROP TABLE IF EXISTS cities CASCADE;
DROP TABLE IF EXISTS users CASCADE;

-- ===============================
-- CREATE TABLES
-- ===============================
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE cities (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE cinemas (
    id BIGSERIAL PRIMARY KEY,
    city_id BIGINT NOT NULL REFERENCES cities(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    address TEXT
);

CREATE TABLE auditoriums (
    id BIGSERIAL PRIMARY KEY,
    cinema_id BIGINT NOT NULL REFERENCES cinemas(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    seat_rows INT,
    seat_columns INT
);

CREATE TABLE seats (
    id BIGSERIAL PRIMARY KEY,
    auditorium_id BIGINT NOT NULL REFERENCES auditoriums(id) ON DELETE CASCADE,
    seat_row VARCHAR(10) NOT NULL,
    seat_number INT NOT NULL,
    status VARCHAR(20) DEFAULT 'active'
);

CREATE UNIQUE INDEX uniq_seat_per_auditorium 
ON seats (auditorium_id, seat_row, seat_number);

CREATE TABLE movies (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    duration_minutes INT NOT NULL,
    rating VARCHAR(20),
    description TEXT
);

CREATE TABLE showtimes (
    id BIGSERIAL PRIMARY KEY,
    movie_id BIGINT NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    auditorium_id BIGINT NOT NULL REFERENCES auditoriums(id) ON DELETE CASCADE,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    status VARCHAR(20) DEFAULT 'active'
);

CREATE TABLE seat_locks (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    showtime_id BIGINT NOT NULL REFERENCES showtimes(id) ON DELETE CASCADE,
    seat_id BIGINT NOT NULL REFERENCES seats(id) ON DELETE CASCADE,
    locked_at TIMESTAMP DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    is_expired BOOLEAN DEFAULT FALSE
);

CREATE UNIQUE INDEX uniq_active_lock 
ON seat_locks (showtime_id, seat_id)
WHERE is_expired = FALSE;

CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    showtime_id BIGINT NOT NULL REFERENCES showtimes(id),
    total_price NUMERIC(12,2) NOT NULL,
    payment_method VARCHAR(50),
    status VARCHAR(30) DEFAULT 'pending', 
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE tickets (
    id BIGSERIAL PRIMARY KEY,
    transaction_id BIGINT NOT NULL REFERENCES transactions(id) ON DELETE CASCADE,
    seat_id BIGINT NOT NULL REFERENCES seats(id),
    showtime_id BIGINT NOT NULL REFERENCES showtimes(id),
    ticket_code VARCHAR(50) UNIQUE NOT NULL,
    status VARCHAR(30) DEFAULT 'active'
);

CREATE UNIQUE INDEX uniq_seat_per_showtime
ON tickets (showtime_id, seat_id)
WHERE status = 'active';

CREATE TABLE refunds (
    id BIGSERIAL PRIMARY KEY,
    ticket_id BIGINT NOT NULL REFERENCES tickets(id),
    user_id BIGINT NOT NULL REFERENCES users(id),
    refund_amount NUMERIC(12,2) NOT NULL,
    refund_type VARCHAR(20) NOT NULL,
    processed_at TIMESTAMP DEFAULT NOW(),
    reason TEXT
);

-- ===============================
-- INSERT DUMMY DATA
-- ===============================
INSERT INTO users (full_name, email, password_hash)
VALUES
('Dimas Akmal', 'dimas@test.com', '$2a$14$rZP1lhpEIovw.TVsMNCT5OKU136tzKvogLGO8vz1hSiUin1rAMrH6');

INSERT INTO cities (name)
VALUES
('Jakarta'),
('Bandung');

INSERT INTO cinemas (city_id, name, address)
VALUES
(1, 'Cinema XXI Jakarta', 'Jl. Sudirman No.1'),
(2, 'Cinema XXI Bandung', 'Jl. Braga No.10');

INSERT INTO auditoriums (cinema_id, name, seat_rows, seat_columns)
VALUES
(1, 'Studio 1', 5, 5),
(1, 'Studio 2', 4, 4),
(2, 'Studio A', 6, 6);

INSERT INTO seats (auditorium_id, seat_row, seat_number)
VALUES
(1,'A',1),(1,'A',2),(1,'A',3),(1,'A',4),(1,'A',5),
(1,'B',1),(1,'B',2),(1,'B',3),(1,'B',4),(1,'B',5),
(1,'C',1),(1,'C',2),(1,'C',3),(1,'C',4),(1,'C',5),
(1,'D',1),(1,'D',2),(1,'D',3),(1,'D',4),(1,'D',5),
(1,'E',1),(1,'E',2),(1,'E',3),(1,'E',4),(1,'E',5);

INSERT INTO movies (title, duration_minutes, rating, description)
VALUES
('Avengers: Endgame', 181, 'PG-13', 'Superhero action film'),
('Inception', 148, 'PG-13', 'Sci-fi thriller');

INSERT INTO showtimes (movie_id, auditorium_id, start_time, end_time)
VALUES
(1, 1, '2025-12-09 14:00:00', '2025-12-09 17:01:00'),
(2, 2, '2025-12-09 18:00:00', '2025-12-09 20:28:00');
