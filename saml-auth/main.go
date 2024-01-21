package main

// Initialize the IDP URL
idpMetadataURL, err := url.Parse("https://samltest.id/saml/idp")
if err != nil {
panic(err)
}
// Initialize metadata of the IDP
idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient,
*idpMetadataURL)
if err != nil {
panic(err)
}

rootURL, err := url.Parse("http://localhost:8000")
if err != nil {
panic(err)
}
// Create SAML SP with the metadata of IDP
samlSP, _ := samlsp.New(samlsp.Options{
URL:            *rootURL,
Key:            <rsa.PrivateKey>,
Certificate:    <Cert>,
IDPMetadata: idpMetadata,
})
app := http.HandlerFunc(hello)
// Use the samlSP to protect your endpoint
http.Handle("/hello", samlSP.RequireAccount(app))
http.Handle("/saml/", samlSP)
http.ListenAndServe(":8000", nil)