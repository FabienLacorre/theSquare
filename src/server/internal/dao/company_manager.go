package dao

// func (b *DataManager) GetCompanyWithName(pattern string) error {
// 	var results []Company

// 	r, err := b.conn.QueryNeo("MATCH (c:Company) WHERE c.name CONTAINS {pattern} RETURN c", map[string]interface{}{
// 		"pattern": pattern,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("cannot query: %v", err)
// 	}
// 	defer r.Close()

// 	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
// 		d := data[0].(graph.Node)

// 		p := Company{
// 			Name:        d.Properties["name"].(string),
// 			Siret:       d.Properties["siret"].(string),
// 			Siren:       d.Properties["siren"].(string),
// 			Description: d.Properties["description"].(string),
// 		}
// 		results = append(results, p)
// 	}

// 	for i := range results {
// 		fmt.Println(results[i].Name)
// 		fmt.Println(results[i].Siret)
// 		fmt.Println(results[i].Siren)
// 		fmt.Println(results[i].Description)
// 	}

// 	return nil
// }

// func (b *DataManager) SetCompany(name string, siret string, siren string, description string) error {
// 	_, err := b.conn.ExecNeo("CREATE (c:Company {name: {name}, siret: {siret}, siren: {siren} , description: {description}})", map[string]interface{}{
// 		"name":        name,
// 		"siret":       siret,
// 		"siren":       siren,
// 		"description": description,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("cannot query: %v", err)
// 	}

// 	return nil
// }
