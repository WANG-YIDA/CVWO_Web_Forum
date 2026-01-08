import PostList from "../components/PostList";
import { Box, IconButton } from "@mui/material";
import AddIcon from "@mui/icons-material/Add";
import React from "react";

const PostListView: React.FC = () => {
    return (
        <div
            style={{
                display: "flex",
                flexDirection: "column",
                justifyContent: "flex-start",
                alignItems: "center",
                minHeight: "100vh",
                textAlign: "center",
                gap: "6rem",
                paddingTop: "4rem",
            }}
        >
            <Box sx={{ width: "100%", maxWidth: 800, margin: "0 auto" }}>
                <Box sx={{ display: "flex", justifyContent: "flex-end", mr: 2.5, mb: -3.5 }}>
                    <IconButton
                        aria-label="add"
                        sx={{
                            width: 40,
                            height: 25,
                            borderRadius: "8px",
                            backgroundColor: "primary.main",
                            color: "#fff",
                            "&:hover": { backgroundColor: "primary.dark" },
                        }}
                    >
                        <AddIcon sx={{ fontSize: 16 }} />
                    </IconButton>
                </Box>
                <PostList />
            </Box>
        </div>
    );
};

export default PostListView;
