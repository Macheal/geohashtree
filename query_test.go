package geohashtree

import (
	"fmt"
	geojson "github.com/paulmach/go.geojson"
	"math/rand"
	"testing"
)

var testtree, _ = OpenGeohashTreeCSV("test_data/a.csv")

//var testtree2, _ = OpenGeohashTreeBoltDB("../county.db")

var testtree2, _ = OpenGeohashTreeBoltDB("test_data/a.db")
var keys = getkeys("test_data/a.csv")

func getkeys(filename string) []string {
	scanner, _ := NewScannerFile(filename)
	keys := []string{}
	for scanner.Next() {
		key, _ := scanner.KeyValue()
		keys = append(keys, key)
	}
	return keys
}

func BenchmarkQueryMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testtree.Query(RandomPt())
	}
}

func BenchmarkQueryBoltDB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testtree2.Query(RandomPt())
	}
}

func BenchmarkGetMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testtree.Get(keys[rand.Intn(len(keys)-1)])
	}
}

func BenchmarkGetBoltDB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testtree2.Get(keys[rand.Intn(len(keys)-1)])
	}
}

func TestGetGeoHashByLocalID(t *testing.T) {
	//geoStr := `{"geometry":{"type":"Polygon","coordinates":[[[116.097078323364,39.9063949275062],[116.095721125603,39.9056912570992],[116.095780134201,39.9048353207799],[116.097319722176,39.9048682416052],[116.097512841225,39.906209751784],[116.097078323364,39.9063949275062]]]}}`
	geoStr := `{"type":"Feature","properties":{"name":"永定政府"},"geometry":{"type":"Polygon","coordinates":[[[116.097078323364,39.9063949275062],[116.095721125603,39.9056912570992],[116.095780134201,39.9048353207799],[116.097319722176,39.9048682416052],[116.097512841225,39.906209751784],[116.097078323364,39.9063949275062]]]}}`
	//geoStr := `{"geometry": {"type": "Polygon", "coordinates": [[[-97.94860839843749, 42.44778143462245], [-97.97607421875, 42.65820178455667], [-98.5308837890625, 42.44372793752476], [-98.8714599609375, 42.06560675405716], [-98.85498046875, 42.459940352216556], [-98.9813232421875, 42.342305278572816], [-98.997802734375, 41.795888098191426], [-98.953857421875, 41.35619553438905], [-98.5968017578125, 41.7180304600481], [-98.624267578125, 41.95949009892467], [-98.173828125, 42.204107493733176], [-97.9705810546875, 42.020732852644294], [-98.349609375, 41.89001042401827], [-98.118896484375, 41.84910468610387], [-97.833251953125, 41.857287927691345], [-97.72338867187499, 42.248851700720955], [-97.22351074218749, 42.49235259142821], [-97.09167480468749, 42.0615286181226], [-97.020263671875, 42.62183364891663], [-97.94860839843749, 42.44778143462245]]]}, "type": "Feature", "properties": {}}`

	//geoStr := `{"type":"Feature","properties":{},"geometry":{"type":"Polygon","coordinates":[[[116.31362915039061,39.99500778093748],[116.22024536132811,39.86758762451019],[116.38092041015625,39.768436410838426],[116.53060913085936,39.86600654754002],[116.4715576171875,39.996585880995035],[116.31362915039061,39.99500778093748]]]}}`
	feature, _ := geojson.UnmarshalFeature([]byte(geoStr))
	ch := MakePolygonIndex2(feature.Geometry.Polygon, 0, 7)
	i := 0
	for {

		data, ok := <-ch
		if !ok {
			break
		}
		i++
		fmt.Println(data)
	}
	fmt.Println("count:", i)
}

func TestA(t *testing.T) {
	list := []string{}
	ExpandGeohashLv("wx4dx", 7, &list)
	fmt.Println(len(list))
	fmt.Println(list)
}

func TestGeohash(t *testing.T) {
	en := Geohash([]float64{116.097078323364, 39.9063949275062}, 7)
	fmt.Println(en)

}
