package constants

type Cors struct {
}

func (*Cors) GetOrigins() []string {
	return []string{
		"http://localhost:3002", //dev
		// Insert cors origins here
	}
}

func (*Cors) GetAllowMethods() []string {
	return []string{"GET", "HEAD", "OPTIONS", "POST", "PUT", "PATCH"}
}

func (*Cors) GetAllowHeaders() []string {
	return []string{"Origin", "Content-Type", "Authorization"}
}

func (*Cors) GetExposedHeaders() []string {
	return []string{"Content-Length"}
}
