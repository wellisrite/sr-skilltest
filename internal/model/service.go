package model

type (
	Properties struct {
		App      App      `json:"app"`
		Database Database `json:"database"`
		Service  Service  `json:"service"`
		Cache    Cache    `json:"cache"`
	}

	App struct {
		Mode            string `json:"mode"`
		Debug           string `json:"debug"`
		CoachPortal     string `json:"coach_portal"`
		BaseURL         string `json:"base_url"`
		DjangoFCM       string `json:"django_fcm"`
		CorsAllowOrigin string `json:"cors_allow_origin"`
	}

	Database struct {
		DBHost     string `json:"db_host"`
		DBName     string `json:"db_name"`
		DBUser     string `json:"db_user"`
		DBPassword string `json:"db_password"`
		DBPort     string `json:"db_port"`
	}

	Cache struct {
		CacheHost     string `json:"cache_host"`
		CachePassword string `json:"cache_password"`
		CacheDB       int    `json:"cache_db"`
		CachePort     string `json:"cache_port"`
	}

	Service struct {
		ServicePort string `json:"service_port"`
		TimeZone    string `json:"time_zone"`
		PoolSize    int    `json:"pool_size"`
		LogPath     string `json:"log_path"`
	}
)
