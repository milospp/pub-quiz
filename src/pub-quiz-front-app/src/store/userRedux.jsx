import { createSlice } from '@reduxjs/toolkit'

export const userRedux = createSlice({
  name: 'userRedux',
  initialState: {
    value: null,
  },
  reducers: {
    setUser: (state, action) => {
      // console.log(action.payload);

      window.sessionStorage.setItem("user", JSON.stringify(action.payload))

      state.value = action.payload;
    },

  },
})

// Action creators are generated for each case reducer function
export const { setUser } = userRedux.actions

export default userRedux.reducer
