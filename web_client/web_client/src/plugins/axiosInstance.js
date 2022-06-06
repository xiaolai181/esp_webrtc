import axios from "axios";

const API =axios.create({
    baseURL:"http://localhost:8080",
    timeout: 2000 
}
)
export default API
