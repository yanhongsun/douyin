package config

import "time"

type DatabaseSettingS struct {
	DBType   string
	UserName string
	Password string
	Host     string
	DBName   string
	// TablePrefix  string
	Charset        string
	ParseTime      bool
	MaxIdleConns   int
	MaxOpenConns   int
	UserTableName  string
	VideoTableName string
	CommTableName  string
	RelaTableName  string
	FavorTableName string
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type ServerSettingS struct {
	EtcdHost     string
	ApiServName  string
	ApiServHost  string
	UserServName string
	UserServHost string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
