import { configureStore } from '@reduxjs/toolkit'
import userRedux from './store/userRedux'

export default configureStore({
  reducer: {
      user: userRedux,
  },
})