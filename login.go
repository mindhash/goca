package goca


func (client *OneClient) UserLogin(username string, token string, timeSeconds int, effectiveGID uint) (string, error) {
  response, err:= client.Call("one.user.login", username, "", -1, 1)
  return response.Body(), err
}