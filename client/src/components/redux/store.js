import { configureStore } from '@reduxjs/toolkit';
import userReducer from './userSlice';
import  taskReducer  from "./taskSlice"
import authReducer from "./authSlice"

export const store = configureStore({
  reducer: {
    user: userReducer,
    task: taskReducer,
    auth: authReducer
  }
});
