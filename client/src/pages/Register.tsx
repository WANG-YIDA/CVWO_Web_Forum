import { Box, Button, Link, Paper, TextField, Typography } from "@mui/material";
import React from "react";

const Register: React.FC = () => {
    return (
        <div>
            <Box display="flex" justifyContent="center" alignItems="center" minHeight="100vh">
                <Paper elevation={3} sx={{ p: 4, maxWidth: 400, width: "100%" }}>
                    <Typography variant="h4" gutterBottom align="center" sx={{ fontWeight: "bold" }}>
                        Sign Up
                    </Typography>

                    <form>
                        <TextField fullWidth label="Username" margin="normal" variant="outlined">
                            Username
                        </TextField>
                    </form>

                    <Button type="submit" variant="contained" sx={{ mt: 2 }}>
                        Create Account
                    </Button>

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
