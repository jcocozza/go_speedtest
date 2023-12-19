-- creates the database table for storing speedtest results
CREATE TABLE IF NOT EXISTS speedtest_results (
    id INT PRIMARY KEY AUTO_INCREMENT,
    `type` VARCHAR(50),
    `timestamp` DATETIME,
    ping_jitter FLOAT,
    ping_latency FLOAT,
    ping_low FLOAT,
    ping_high FLOAT,
    download_bandwidth INT,
    download_bytes INT,
    download_elapsed INT,
    download_latency_iqm FLOAT,
    download_latency_low FLOAT,
    download_latency_high FLOAT,
    download_latency_jitter FLOAT,
    upload_bandwidth INT,
    upload_bytes INT,
    upload_elapsed INT,
    upload_latency_iqm FLOAT,
    upload_latency_low FLOAT,
    upload_latency_high FLOAT,
    upload_latency_jitter FLOAT,
    packetLoss FLOAT,
    isp VARCHAR(100),
    interface_internalIp VARCHAR(50),
    interface_name VARCHAR(50),
    interface_macAddr VARCHAR(50),
    interface_isVpn BOOLEAN,
    interface_externalIp VARCHAR(50),
    server_id INT,
    server_host VARCHAR(100),
    server_port INT,
    server_name VARCHAR(100),
    server_location VARCHAR(100),
    server_country VARCHAR(100),
    server_ip VARCHAR(50),
    result_id VARCHAR(100),
    result_url VARCHAR(255),
    result_persisted BOOLEAN
);
-- creates a database table for storing data associated with the weather during a given speedtest
CREATE TABLE IF NOT EXISTS weather (
    id INT PRIMARY KEY AUTO_INCREMENT,
    SpeedTest_id INT
    FOREIGN KEY (SpeedTest_id) REFERENCES speedtest_results(id)

     -- Location Fields
    LocationName TEXT,
    Region TEXT,
    Country TEXT,
    Lat FLOAT,
    Lon FLOAT,
    TzID TEXT,
    LocaltimeEpoch INTEGER,
    Localtime DATETIME,

    -- Current Fields
    LastUpdatedEpoch INT,
    LastUpdated DATETIME,
    TempC FLOAT,
    TempF FLOAT,
    IsDay BOOLEAN,

    -- Condition Fields
    ConditionText TEXT,
    ConditionIcon TEXT,
    ConditionCode INTEGER,

    -- Current
    -- Wind Fields
    WindMph FLOAT,
    WindKph FLOAT,
    WindDegree INT,
    WindDir TEXT,
    -- Pressure Fields
    PressureMb FLOAT,
    PressureIn FLOAT,
    -- Precipitation Fields
    PrecipMm FLOAT,
    PrecipIn FLOAT,
    -- Additional Weather Fields
    Humidity INTEGER,
    Cloud INTEGER,
    FeelslikeC FLOAT,
    FeelslikeF FLOAT,
    VisKm FLOAT,
    VisMiles FLOAT,
    Uv FLOAT,
    GustMph FLOAT,
    GustKph FLOAT,
    -- Air Quality Fields
    Co FLOAT,
    No2 FLOAT,
    O3 FLOAT,
    So2 FLOAT,
    Pm25 FLOAT,
    Pm10 FLOAT,
    UsEpaIndex INTEGER,
    GbDefraIndex INTEGER NOT NULL
);