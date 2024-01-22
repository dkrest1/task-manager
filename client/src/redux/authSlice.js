/* eslint-disable no-unused-vars */
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import { signup, login } from '../api/api';

const initialState = {
  loading: false,
  success: false, 
  userInfo: {}, 
  error: null,
  token: null,
  isAuthenticated: false
  
}

// Signup
export const signupAsync = createAsyncThunk("auth/signup", async(userData, {rejectWithValue}) => {
  try{
    const response = await signup(userData)
    return response
  }catch(error) {
    return rejectWithValue(error)
  }
})

// Login 
export const loginAsync = createAsyncThunk("auth/login", async(userData, {rejectWithValue}) => {
  try {
    const response = await login(userData)
    return response
  }catch(error) {
    return rejectWithValue(error)
  }
})

// Set Error
export const setError = (errorMessage) => ({ type: 'auth/setError', payload: errorMessage });

const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    logout: (state) => {
      state.isAuthenticated = false;
      state.userInfo = {};
      state.token = null;
      state.success = false
      state.loading = false
      state.success = false 
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(signupAsync.pending, (state) => {
        state.loading = true
      })
      .addCase(signupAsync.fulfilled, (state, action) => {
        state.loading  = false
        state.success = true
        state.userInfo = action.payload.payload

      })
      .addCase(signupAsync.rejected, (state, action) => {
        state.loading = false
        const errorPayload = action.payload;

        if (errorPayload.status === 400) {
          state.error = errorPayload.message;
        } else {
          state.error = 'Signup failed. Please try again.';
        }
      })
      .addCase(loginAsync.pending, (state) => {
        state.loading = true
      })
      .addCase(loginAsync.fulfilled, (state, action) => {
        state.loading = false
        state.success = true
        state.userInfo = action.payload.payload.user
        state.isAuthenticated = true
        state.token = action.payload.payload.token
        state.error = false
      })
      .addCase(loginAsync.rejected, (state, action) => {
        state.loading = false
        state.success = false
        state.isAuthenticated = false
        state.userInfo = {}
        state.token = null
        const errorPayload = action.payload
        
        if(errorPayload.status === 400) {
          state.error = errorPayload.message
        }else if(errorPayload.status === 500) {
          state.error = errorPayload.message
        }else {
          state.error = 'Login failed. Please try again.';
        }
      })
  }
});

export const { logout } = authSlice.actions;

export default authSlice.reducer;

