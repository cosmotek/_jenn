package registry

// type Registry struct {
// 	enums map[string][]string
// }

// func (r *Registry) Add(enumName string, values ...string) {
// 	if r.enums == nil {
// 		r.enums = map[string][]string{}
// 	}

// 	r.enums[enumName] = values
// }

// func (r *Registry) FindEnum(typeName string) (primitives.Enum, bool) {
// 	val, ok := r.enums[typeName]
// 	if !ok {
// 		return primitives.Enum{}, false
// 	}

// 	return primitives.Enum{
// 		Name:   typeName,
// 		Values: val,
// 	}, true
// }
