package cache

import "time"


type item struct {
    value      string
    deadline int64
}

type Cache struct {
	myMap map[string]*item
}

func NewCache() Cache {
	cache:= Cache{
		myMap: make(map[string]*item),
	}
	go func() {
        for now := range time.Tick(time.Second) {
           
            for k, v := range cache.myMap {
				if(v.deadline !=0){
                if now.Unix() - v.deadline >= 0 {
                    delete(cache.myMap, k)
                }
			}
            }
    
        }
    }()
	return  cache
}

func (cache *Cache) Get(key string) (string, bool) {
	var value string;
    if i, ok := cache.myMap[key]; ok {
        return i.value,ok
     }
     return value,false

}

func (cache *Cache) Put(key, value string) {
	  i := &item{value: value}  
      cache.myMap[key] = i	  
}

func (cache *Cache) Keys() []string {
	keys := make([]string, 0, len(cache.myMap))
	for k := range cache.myMap {
		keys = append(keys, k)
	}
	return keys
}


func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	  i := &item{value: value}  
      cache.myMap[key] = i	
	  i.deadline=time.Time(deadline).Unix()
}