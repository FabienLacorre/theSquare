package dao

// func (b *DataManager) GetJobWithName(pattern string) error {
// 	var results []Job

// 	r, err := b.conn.QueryNeo("MATCH (j:Job) WHERE j.name CONTAINS {pattern} RETURN j", map[string]interface{}{
// 		"pattern": pattern,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("cannot query: %v", err)
// 	}
// 	defer r.Close()

// 	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
// 		d := data[0].(graph.Node)

// 		p := Job{
// 			Name:        d.Properties["name"].(string),
// 			Description: d.Properties["description"].(string),
// 			GrossWage:   d.Properties["grossWage"].(string),
// 		}
// 		results = append(results, p)
// 	}

// 	for i := range results {
// 		fmt.Println(results[i].Name)
// 		fmt.Println(results[i].Description)
// 		fmt.Println(results[i].GrossWage)
// 	}

// 	return nil
// }

// func (b *DataManager) SetJob(name string, description string, grossWage string) error {
// 	_, err := b.conn.ExecNeo("CREATE (j:Job {name: {name}, description: {description}, grossWage: {grossWage}})", map[string]interface{}{
// 		"name":        name,
// 		"description": description,
// 		"grossWage":   grossWage,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("cannot query: %v", err)
// 	}

// 	return nil
// }
