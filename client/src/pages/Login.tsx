import { Box, Button, Link, Paper, TextField, Typography } from "@mui/material";
import React from "react";

const Login: React.FC = () => {
    return (
        <div
            style={{
                backgroundImage: "url('/images/home_background.jpg')",
                backgroundSize: "cover",
                backgroundPosition: "center",
                backgroundRepeat: "no-repeat",
            }}
        >
            <Box display="flex" justifyContent="center" alignItems="center" minHeight="100vh">
                <Paper elevation={3} sx={{ p: 4, maxWidth: 400, width: "100%" }}>
                    <Typography variant="h4" gutterBottom align="center">
                        Login
                    </Typography>

                    <form>
                        <TextField fullWidth label="Username" margin="normal" variant="outlined">
                            Username
                        </TextField>
                    </form>

                    <Button type="submit" variant="contained" sx={{ mt: 2 }}>
                        Login
                    </Button>

                    <Typography variant="body2" sx={{ mt: 2 }}>
                        {`Don't have an account? `}
                        <Link href="/register">Sign up</Link>
                    </Typography>
                </Paper>
            </Box>
        </div>
    );
};

export default Login;
