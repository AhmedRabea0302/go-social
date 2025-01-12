import { useNavigate, useParams } from "react-router-dom"
import { API_URL } from "./App"


export const ConfirmationPage = () => {

    const params = useParams()
    const redirect = useNavigate()
    
    const handleConfirm = async () => {
        const token = params.token
        const response = await fetch(`${API_URL}/users/activate/${token}`,{
            method: "PUT"
        })

        if (response.ok) {
            redirect("/")
        } else {
            // handle the error
        }
    }

    return (
        <div>
            <h1>Confirmation</h1>
            <button onClick={handleConfirm}>Click to confirm</button>
        </div>
    )
}