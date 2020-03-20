package goca


// User represents an OpenNebula User
type User struct {
  XMLResource
  ID   uint
  Name string
  Client *OneClient

}

func (client *OneClient) NewUser(id uint) *User {
  return &User{ID: id, Client: client}
}

func (user *User) Login(token string, timeSeconds int, effectiveGID uint) error {
  _, err := user.Client.Call("one.user.login", user.ID, token, timeSeconds, int(effectiveGID))
  return err
}
