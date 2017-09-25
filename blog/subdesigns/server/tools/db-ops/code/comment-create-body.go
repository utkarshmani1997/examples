


	var comment types.Comment
	var iface interface{}
	iface = &comment
	// unmarshal data file into struct
	_, cerr := io.ReadFile(dataFile, &iface)
	if cerr != nil {
		fmt.Println("Error", cerr)
		os.Exit(1)
	}

	fmt.Printf("%+v\n\n", comment)

	// validate the input here

	
	err := types.CreateComment(
		postgres.POSTGRES,
		&comment,
		postUUID,
		userUUID,
	)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println("Created:", comment.UUID)
	



// HOFSTADTER_BELOW

