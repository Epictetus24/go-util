package box

func ChkDom(domain string) {
	var hope []byte
	if domain == "" {
		return
	}

	hope = domain
	host, _ := os.Hostname()
	dnsenv := os.Getenv("USERDNSDOMAIN")
	if !strings.Contains(host, string(hope)) && !strings.Contains(dnsenv, string(hope)) {
		time.Sleep({{.Delay}} * time.Seconds)
		os.Exit(0)
	}

}
