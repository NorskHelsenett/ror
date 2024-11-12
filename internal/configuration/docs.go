// the configuration package is responsible for loading and merging configurations
// it implements loaders and parsers
//
// Example:
//
//	    teststring := `{"birdType":"Pigeon","what it does":"likes to eat seed", "age":12}`
//	    teststring2 := fmt.Sprintf("birdType: %s\nwhat it does: likes to eat seed\nage: %d\n", "Pigeon", 14)
//
//	    generator := configuration.NewConfigurationsGenerator()
//
//	    layer := configuration.NewConfigurationLayer("test1", 2, 250, configuration.NewStringLoader(configuration.ParserTypeJson, teststring))
//	    generator.AddConfiguration(layer)
//
//	    layer2 := configuration.NewConfigurationLayer("test2", 1, 251, configuration.NewStringLoader(configuration.ParserTypeYaml, teststring2))
//	    generator.AddConfiguration(layer2)
//
//	    generator.AddConfiguration(configuration.NewConfigurationLayer("test3", 3, 252, configuration.NewFuncLoader(configuration.ParserTypeYaml, testloaderz)))
//
//		    ownerref := apiresourcecontracts.ResourceOwnerReference{
//			   Scope:   "cluster",
//		       Subject: "sdi-ror-dev-b40y",
//	     }
//
//	     generator.AddConfiguration(configuration.NewConfigurationLayer("test4", 3, 252, configuration.NewFileLoader(configuration.ParserTypeYaml, "../../assets/configs.yaml")))
//	     generator.AddConfiguration(configuration.NewConfigurationLayer("test5", 3, 254, configuration.NewFileLoader(configuration.ParserTypeJson, "../../assets/config1.json")))
//
//	     generator.AddConfiguration(configuration.NewConfigurationLayer("test6", 3, 253, configuration.NewResourceLoader(ownerref, "test")))
//
//	     merged, err := generator.GenerateConfig()
//
//	   	if err != nil {
//		    	fmt.Println(err)
//			}
//
//		    fmt.Println(string(merged))
package configuration
