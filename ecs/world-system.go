package ecs


//func (w *World) Update() {
//	for _, system := range w.systems {
//		// if system.entitiesSystemsDirty.length {
//		// 	this.cleanDirtyEntities();
//		// }
//		system.UpdateAll()
//	}
//}
//
//func (w *World) AddSystem(s System) {
//	s.Init(name, w, s)
//	w.systems = append(w.systems, s)
//	for _, entity := range w.entities {
//		if s.Test(entity) {
//			s.AddEntity(entity)
//		}
//	}
//}
//
//func (w *World) RemoveSystem(s BaseSystem) {
//	// todo
//}
//
//func (w *World) cleanDirtyEntities() {
//	for _, entity := range w.systemsDirtyEntities {
//		for _, system := range w.systems {
//			// entity.systems
//			entityTest := system.Test(entity)
//			if entityTest {
//				system.AddEntity(entity)
//			} else {
//				system.RemoveEntity(entity)
//			}
//		}
//		entity.systemsDirty = false
//	}
//	w.systemsDirtyEntities = nil
//}
