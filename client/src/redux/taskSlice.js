/* eslint-disable no-unused-vars */
import { createSlice } from '@reduxjs/toolkit';

const taskSlice = createSlice({
  name: 'task',
  initialState: {
    tasks: [],
  },
  reducers: {
    setTasks: (state, action) => {
      state.tasks = action.payload;
    },
    addTask: (state, action) => {
      state.tasks.push(action.payload);
    },
    updateTask: (state, action) => {
      // Update task logic
    },
    deleteTask: (state, action) => {
      // Delete task logic
    },
  },
});

export const { addTask, removeTask, updateTask, deleteTask } = taskSlice.actions;
export default taskSlice.reducer;