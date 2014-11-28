/*h := http.FileServer(http.Dir("/home/qboxtest/Desktop"))
	http.ListenAndServe(":3808", h)
这种可能会有延迟,用httptest比较好
	*/
ico, err := ioutil.ReadFile("../../watermark/favicon.ico")
if err != nil {
ts.Fatal(t, "Read file failed:", err)
}
httpLocal := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
w.Write(ico)
}))
println("xxxxxx", httpLocal.URL)

httpLocal := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
http.ServeFile(w, r, fileName)
}))