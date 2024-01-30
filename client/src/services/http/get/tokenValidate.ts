import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:800',
  withCredentials: true, 
});

api.interceptors.request.use(
  (config) => {
    return config;
  },
  (error) => {
    
    return Promise.reject(error);
  }
);

api.interceptors.response.use(
  (response) => {
    return response;
  },
  async (error) => {
    if (error.response && error.response.status === 401) {
     
      try {

        // If the token is refreshed successfully, retry the original request
        return api.request(error.config);
      } catch (refreshError) {
        console.error('Token refresh failed:', refreshError);
        // Optionally, redirect the user to the login page or show an error message

      }
    }

    return Promise.reject(error);
  }
);

export default api;
