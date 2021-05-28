goodCities := groupCity[100]
for city, _ := range cityPopulation {
    cityIndex := -1
    for i, curr := range goodCities {
        if curr == city {
            cityIndex = i
            break
        }
    }
    if cityIndex == -1 {
        delete(cityPopulation, city)
    }
}