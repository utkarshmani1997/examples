


	var post types.Post
	var iface interface{}
	iface = &post
	// unmarshal data file into struct
	_, cerr := io.ReadFile(dataFile, &iface)
	if cerr != nil {
		fmt.Println("Error", cerr)
		os.Exit(1)
	}

	fmt.Printf("%+v\n\n", post)

	// validate the input here

	
	err := types.CreatePost(
		postgres.POSTGRES,
		&post,
		userUUID,
	)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println("Created:", post.UUID)
	



// HOFSTADTER_BELOW

