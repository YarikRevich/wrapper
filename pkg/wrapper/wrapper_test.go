package wrapper

import (
	"encoding/json"
	"testing"

	"github.com/franela/goblin"
)

func TestWrapper(t *testing.T){
	g := goblin.Goblin(t)
	g.Describe("TestWrapper", func ()  {
		type Stub struct {
			Name string
			Age int
		}

		g.It("Test set and get base", func(){
			
			s := Stub{Name: "Yarik", Age: 10}
			wrapper := UseWrapper()
			wrapper.SetBase(s)

			g.Assert(wrapper.GetBase()).Eql(s)
		})

		g.It("Test set field", func(){
			
			s := Stub{Name: "Yarik", Age: 10}

			wrapper := UseWrapper()
			wrapper.SetBase(s)
			wrapper.SetField("id", 1)
			g.Assert(wrapper.GetField("id")).Eql(1)
		})

		g.It("Test marshaler", func(){
			
			s := Stub{Name: "Yarik", Age: 10}

			wrapper := UseWrapper()
			wrapper.SetEncoder(json.Marshal)
			
			wrapper.SetBase(s)
			wrapper.SetField("id", 1)
			
			b, err := wrapper.Marshal()
			g.Assert(err).IsNil()
			g.Assert(string(b)).Eql("{\"base\":{\"Name\":\"Yarik\",\"Age\":10},\"id\":1}")
		})

		g.It("Test unmarshaler", func(){
			
			s := Stub{Name: "Yarik", Age: 10}

			wrapper := UseWrapper()
			wrapper.SetDecoder(json.Unmarshal)
			
			wrapper.SetBase(s)
			wrapper.SetField("id", 1)
			

			g.Assert(wrapper.Unmarshal([]byte("{\"base\": 10, \"id\": 2}"))).IsNil()
			g.Assert(wrapper.GetBase().(float64)).Eql(float64(10))
			g.Assert(wrapper.GetField("id").(float64)).Eql(float64(2))
		})
	})
}