-- Create tables
CREATE TABLE IF NOT EXISTS campaigns (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cid VARCHAR(255) NOT NULL,
    cname VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL,
    cta VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    update_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_cid (cid),
    INDEX idx_cname (cname),
    INDEX idx_status (status)
);

CREATE TABLE IF NOT EXISTS campaigns_rule (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cid VARCHAR(255) NOT NULL,
    include_os VARCHAR(255) NOT NULL,
    include_country VARCHAR(255) NOT NULL,
    include_app VARCHAR(255) NOT NULL,
    exclude_os VARCHAR(255) NOT NULL,
    exclude_country VARCHAR(255) NOT NULL,
    exclude_app VARCHAR(255) NOT NULL,
    created_at datetime DEFAULT CURRENT_TIMESTAMP,
    update_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_ios (include_os),
    INDEX idx_ic (include_country),
    INDEX idx_ia (include_app),
    INDEX idx_eo (exclude_os),
    INDEX idx_ec (exclude_country),
    INDEX idx_ea (exclude_app)
);

-- Insert sample data
INSERT INTO campaigns (cid, cname, image, cta, status) VALUES
('spotify', 'Spotify - Music for everyone', 'https://cdn.example.com/Spotify.jpg', 'Play', 'ACTIVE'),
('duolingo', 'Sports App Promotion', 'https://cdn.example.com/duolingo.jpg', 'Download', 'ACTIVE'),
('subwaysurfer', 'Subway Surfer', 'https://cdn.example.com/subwaysurfer.jpg', 'Install', 'ACTIVE'),
('templerun', 'Temple Run', 'https://cdn.example.com/templerun.jpg', 'Install', 'ACTIVE'),
('clashofclans', 'Clash Of Clans', 'https://cdn.example.com/clashofclans.jpg', 'Play', 'INACTIVE');

-- Insert data into campaigns_rule table
INSERT INTO campaigns_rule (cid, include_os, include_country, include_app, exclude_os, exclude_country, exclude_app) VALUES
(
    'spotify',
    'android,ios',
    'US,UK,IN,CA',
    'music_app1,music_app2',
    'windows',
    'CN,RU',
    'game_app_1'
),
(
    'duolingo',
    'android,ios,windows',
    'US,UK,AU,NZ',
    'learning_app1,learning_app2',
    'linux',
    'BR,AR',
    'music_app1'
),
(
    'subwaysurfer',
    'ios',
    'IN,SG,MY,PH',
    'game_app1,game_app2',
    'android',
    'CN',
    'music_app1'
),
(
    'templerun',
    'android',
    'US,CA,UK,AU',
    'game_app1,game_app2',
    'ios',
    'RU',
    'music_app2'
),
(
    'clashofclans',
    'android,ios',
    'US,UK,CA,AU,NZ',
    'game_app_1,game_app_2,game_app_3',
    'windows',
    'CN,RU,BR',
    'learning_app1,'
);
