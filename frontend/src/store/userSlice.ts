import { createSlice, PayloadAction } from '@reduxjs/toolkit'

export interface UserState {
  loading: boolean
  username: string
}

const initialState: UserState = {
  loading: false,
  username: '',
}

export const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    startLogin(state) {
      state.loading = true
    },
    loginSuccess(state, action: PayloadAction<{ username: string }>) {
      state.loading = false
      state.username = action.payload.username
    },
  },
})

// Action creators are generated for each case reducer function
export const { startLogin, loginSuccess } = userSlice.actions

export default userSlice.reducer
