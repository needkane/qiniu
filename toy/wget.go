if _, err := os.Stat("ffmpeg"); err != nil {
	        f, err := os.Create("ffmpeg")
			        err = os.Chmod("ffmpeg", 0777)
					        fmt.Println("ffmpeg chmod failed", err)
							if err != nil {
								            fmt.Println("os create file failed")
											        
							}   
							        //var ffmpegUrl = "http://s3-us-west-2.amazonaws.com/qiniu-bs/2014-11-18/ffmpeg"
									        var ffmpegUrl = "http://7xlywb.com1.z1.glb.clouddn.com/uffmpeg"
											if os.Getenv("$UNAME") == "Darwin" {
												            ffmpegUrl = "http://7xlywb.com1.z1.glb.clouddn.com/macFfmpeg"
															        
											}   
											        resp, err := http.Get(ffmpegUrl)
													        defer resp.Body.Close()
															if err != nil {
																            fmt.Println("http get ffmpeg failed")
																			        
															}   
															        io.Copy(f, resp.Body)
																	 
																	        fmt.Println("download success")
																			        time.Sleep(5 * time.Second)
																					    
}   

