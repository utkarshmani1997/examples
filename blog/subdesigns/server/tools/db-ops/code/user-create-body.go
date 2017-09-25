


	var user types.User
	var iface interface{}
	iface = &user
	// unmarshal data file into struct
	_, cerr := io.ReadFile(dataFile, &iface)
	if cerr != nil {
		fmt.Println("Error", cerr)
		os.Exit(1)
	}

	fmt.Printf("%+v\n\n", user)

	// validate the input here

	
	err := types.CreateUser(
		postgres.POSTGRES,
		&user,
	)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	fmt.Println("Created:", user.UUID)
	



// HOFSTADTER_BELOW

