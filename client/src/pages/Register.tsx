import { Box, Button, Link, Paper, TextField, Typography } from "@mui/material";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const Register: React.FC = () => {
    const [username, setUsername] = useState("");
    const [error, setError] = useState("");
    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent) => {
        // prevent page reload
        e.preventDefault();

        // clear previous error
        setError("");

        //fetch backend for validation and existence check
        try {
            // Request to POST register info
            const response = await fetch("http://localhost:8000/api/auth/register", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ username: username }),
            });
            const data_json = await response.json();

            // process data to get result
            if (data_json.success && data_json.payload?.data) {
                const auth_result = data_json.payload.data;

                if (auth_result.success && auth_result.user) {
                    const user_id = auth_result.user.id;

                    // keep track of user id for current seesion
                    sessionStorage.setItem("user_id", user_id.toString());

                    // navigate to TopicListView page
                    navigate("/topics");
                } else {
                    // print client-oriented error message
                    const error_message = auth_result.error;
                    setError(error_message);
                }
            } else {
                console.error("Failed to POST register info: %s", data_json.error);
                setError("Server error :(");
            }
        } catch (error) {
            console.error("Error fetching register:", error);
            setError("Network error :(");
        }
    };

    return (
        <div>
            <Box display="flex" justifyContent="center" alignItems="center" minHeight="100vh">
                <Paper elevation={3} sx={{ p: 4, maxWidth: 400, width: "100%" }}>
                    <Typography variant="h4" gutterBottom align="center" sx={{ fontWeight: "bold" }}>
                        Sign Up
                    </Typography>

                    <form onSubmit={handleSubmit}>
                        <TextField
                            fullWidth
                            label="Username"
                            margin="normal"
                            variant="outlined"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                        />

                        {error && (
                            <Typography color="error" variant="body2" sx={{ mt: 1 }}>
                                {error}
                            </Typography>
                        )}

                        <Button type="submit" variant="contained" sx={{ mt: 2 }}>
                            Create Account
                        </Button>
                    </form>

                    <Typography variant="body2" sx={{ mt: 2 }}>
                        {`Already have an account? `}
                        <Link href="/login">Login</Link>
                    </Typography>
                </Paper>
            </Box>
        </div>
    );
};

export default Register;
