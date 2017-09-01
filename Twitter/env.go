package Twitter

type TwitterConfig struct {
	ConfKey     string
	ConfSecret  string
	TokenKey    string
	TokenSecret string
}

// There were too many of cof keys, and my env var were messed.
// So I used This simple method, BUT you could use os.Args() method instead.
// AND! Please, do not POST this on web.
func conf(env TwitterConfig) TwitterConfig {

	env.ConfKey = ""
	env.ConfSecret = ""
	env.TokenKey = ""
	env.TokenSecret = ""

	return env
}
