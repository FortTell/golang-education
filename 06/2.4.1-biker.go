type MadMaxRoboBiker struct {
    On bool
    Ammo int
    Power int
}

func (biker *MadMaxRoboBiker) Shoot() bool {
    if !biker.On || biker.Ammo <= 0 {
        return false
    }
    
    biker.Ammo--
    return true
}

func (biker *MadMaxRoboBiker) RideBike() bool {
    if !biker.On || biker.Power <= 0 {
        return false
    }
    
    biker.Power--
    return true
}
