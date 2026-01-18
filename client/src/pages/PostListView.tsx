import Post from "../types/Post";
import PostList from "../components/PostList";
import AddIcon from "@mui/icons-material/Add";
import {
    Alert,
    Box,
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    IconButton,
    Snackbar,
    TextField,
    Typography,
} from "@mui/material";
import { useNavigate, useParams } from "react-router-dom";
import React, { useEffect, useState } from "react";

interface PostJSON {
    id: number;
    topic_id: number;
    title: string;
    content: string;
    user_id: number;
    author: string;
    created_at: string;
}

const API_DOMAIN = process.env.REACT_APP_API_DOMAIN;
const API_PORT = process.env.REACT_APP_API_PORT;
const API_URL = `${API_DOMAIN}:${API_PORT}`;

const PostListView: React.FC = () => {
    const [topicName, setTopicName] = useState("");
    const [posts, setPosts] = useState<Post[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [open, setOpen] = useState(false);
    const [error, setError] = useState("");
    const [userID, setUserID] = useState<number | null>(null);
    const [showSuccess, setShowSuccess] = useState(false);
    const [showDeleteSuccess, setShowDeleteSuccess] = useState(false);
    const [server_error, setServerError] = useState(false);
    const navigate = useNavigate();
    const { topicID } = useParams<{ topicID: string }>(); // Get topicID from URL params

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        // prevent page reload
        event.preventDefault();

        // clear previous error
        setError("");

        // get form data
        const formData = new FormData(event.currentTarget);
        const formJson = Object.fromEntries(formData.entries() as Iterable<[string, FormDataEntryValue]>);
        const post_title = formJson["post-title"];
        const post_content = formJson["post-content"];

        try {
            // Request to POST posts
            const response = await fetch(`${API_URL}/api/topics/` + topicID + "/posts", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ title: post_title, content: post_content, user_id: userID }),
            });
            const data_json = await response.json();

            // process data to get result
            if (data_json.success && data_json.payload?.data) {
                const post_result = data_json.payload.data;

                if (post_result.success && post_result.post) {
                    const new_post: Post = {
                        id: post_result.post.id,
                        title: post_result.post.title,
                        topic_id: post_result.topic_id,
                        user_id: post_result.post.user_id,
                        content: post_result.post.content,
                        author: post_result.post.author,
                        timestamp: new Date(post_result.post.created_at),
                    };

                    // put new post at top, close dialog then show success message
                    setPosts((posts) => [new_post, ...posts]);
                    setOpen(false);
                    setShowSuccess(true);
                } else {
                    // print client-oriented error message
                    const error_message = post_result.error;
                    setError(error_message);
                }
            } else {
                console.error("Failed to POST posts: %s", data_json.error);
                setError("Server error :(");
            }
        } catch (error) {
            console.error("Error fetching posts:", error);
            setError("Network error :(");
        }
    };

    // navigate to login page if not login yet, or to topics page if topic id is not valid
    useEffect(() => {
        const user_id = sessionStorage.getItem("user_id");
        if (!user_id) {
            navigate("/login");
        } else if (!topicID) {
            navigate("/topics");
        } else {
            setUserID(parseInt(user_id, 10));
        }
    }, [navigate]);

    useEffect(() => {
        const fetchPosts = async () => {
            try {
                // Request to GET posts
                const response = await fetch(`${API_URL}/api/topics/` + topicID + "/posts", {
                    method: "GET",
                    headers: { "Content-Type": "application/json" },
                });
                const data_json = await response.json();

                // process data to get posts
                if (data_json.success && data_json.payload?.data) {
                    const postListResult = data_json.payload.data;

                    if (postListResult.success && postListResult.posts) {
                        const postList: Post[] = postListResult.posts.map((post: PostJSON) => ({
                            id: post.id,
                            title: post.title,
                            topic_id: post.topic_id,
                            content: post.content,
                            user_id: post.user_id,
                            author: post.author,
                            timestamp: new Date(post.created_at),
                        }));

                        setPosts(postList.reverse());
                    }
                } else {
                    console.error("Failed to GET posts: %s", data_json.error);
                    setServerError(true);
                }
            } catch (error) {
                console.error("Error fetching posts:", error);
                setServerError(true);
            } finally {
                setLoading(false);
            }
        };

        // get topic name
        const fetchTopic = async () => {
            try {
                // Request to GET topic
                const response = await fetch(`${API_URL}/api/topics/` + topicID, {
                    method: "GET",
                    headers: { "Content-Type": "application/json" },
                });
                const data_json = await response.json();

                // process data to get posts
                if (data_json.success && data_json.payload?.data) {
                    const topicResult = data_json.payload.data;

                    if (topicResult.success && topicResult.topic) {
                        const topic_name = topicResult.topic.name;
                        setTopicName(topic_name);
                    }
                } else {
                    console.error("Failred to GET topic: %s", data_json.error);
                    setServerError(true);
                }
            } catch (error) {
                console.error("Error fetching topic:", error);
                setServerError(true);
            }
        };

        fetchTopic();
        fetchPosts();
    }, []);

    return server_error ? (
        <Box sx={{ width: "100%", mt: 8, display: "flex", justifyContent: "center" }}>
            <Alert severity="error" variant="filled">
                Oops! Something went wrong. Please try again later.
            </Alert>
        </Box>
    ) : (
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
                {loading ? (
                    <Typography sx={{ mt: 4, color: "text.secondary" }}>Loading topics...</Typography>
                ) : (
                    <>
                        <Box sx={{ display: "flex", justifyContent: "flex-end", mr: 2.5, mb: -2 }}>
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
                                onClick={handleClickOpen}
                            >
                                <AddIcon sx={{ fontSize: 16 }} />
                            </IconButton>
                        </Box>

                        <Snackbar
                            open={showSuccess}
                            anchorOrigin={{ vertical: "top", horizontal: "center" }}
                            autoHideDuration={3000}
                            onClose={() => setShowSuccess(false)}
                        >
                            <Alert
                                onClose={() => setShowSuccess(false)}
                                severity="success"
                                variant="filled"
                                sx={{ width: "100%" }}
                            >
                                New Post Added!
                            </Alert>
                        </Snackbar>

                        <Snackbar
                            open={showDeleteSuccess}
                            anchorOrigin={{ vertical: "top", horizontal: "center" }}
                            autoHideDuration={3000}
                            onClose={() => setShowDeleteSuccess(false)}
                        >
                            <Alert
                                onClose={() => setShowDeleteSuccess(false)}
                                severity="success"
                                variant="filled"
                                sx={{ width: "100%" }}
                            >
                                Post Deleted!
                            </Alert>
                        </Snackbar>

                        <PostList posts={posts} topicID={topicID || ""} topicName={topicName} />
                    </>
                )}
            </Box>

            <Dialog
                open={open}
                onClose={handleClose}
                maxWidth="md"
                fullWidth
                sx={{ "& .MuiDialog-paper": { width: "100%", maxWidth: 900, height: "100%", maxHeight: 500 } }}
            >
                <DialogTitle>Create New Post</DialogTitle>
                <DialogContent>
                    <DialogContentText>Share your thoughts and spark a new discussion!</DialogContentText>
                    <form onSubmit={handleSubmit} id="create-post-form" style={{ paddingTop: 2 }}>
                        <TextField
                            autoFocus
                            required
                            margin="dense"
                            id="post-title"
                            name="post-title"
                            label="Title"
                            type="text"
                            fullWidth
                            variant="standard"
                        />
                        <TextField
                            margin="dense"
                            id="post-content"
                            name="post-content"
                            label="Content"
                            type="text"
                            fullWidth
                            multiline
                            InputLabelProps={{ shrink: true }}
                            minRows={8}
                            sx={{ mt: 4 }}
                        />
                    </form>
                </DialogContent>

                {error && (
                    <Typography color="error" variant="body2" sx={{ mt: 1, textAlign: "center" }}>
                        {error}
                    </Typography>
                )}

                <DialogActions>
                    <Button onClick={handleClose}>Cancel</Button>
                    <Button type="submit" form="create-post-form">
                        Create
                    </Button>
                </DialogActions>
            </Dialog>
        </div>
    );
};

export default PostListView;
