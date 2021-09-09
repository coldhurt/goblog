import { useDispatch, useSelector } from 'react-redux'
import { Post } from '../utils/fetch'
import { RootState } from '../store'
import { startLogin } from '../store/userSlice'

const useLogin = () => {
  const user = useSelector((state: RootState) => state.user)
  const dispatch = useDispatch()
  const login = (username = '', password = '') => {
    dispatch(startLogin())
    Post('admin/login', {
      body: {
        username,
        password,
      },
    })
      .then((res) => res.json())
      .then((res) => {
        console.log(res)
      })
  }
  return {
    user,
    login,
  }
}

export default useLogin
