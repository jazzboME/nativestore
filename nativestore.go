package nativestore

import dcred "github.com/docker/docker-credential-helpers/credentials"

// Set inserts the credential into the system native key store
func Set(lbl, url, user, secret string) error {
	cr := &dcred.Credentials{
		ServerURL: url,
		Username:  user,
		Secret:    secret,
	}

	dcred.SetCredsLabel(lbl)
	return ns.Add(cr)
}

// Get returns previously stored credentials
func Get(lbl, url string) (string, string, error) {
	dcred.SetCredsLabel(lbl)
	return ns.Get(url)
}

// Del removes the credential from the system
func Del(lbl, url string) error {
	dcred.SetCredsLabel(lbl)
	return ns.Delete(url)
}
