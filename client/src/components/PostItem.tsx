import Post from "../types/Post";
import { makeStyles } from "@mui/styles";
import {
    Alert,
    Box,
    Button,
    Card,
    CardActions,
    CardContent,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    IconButton,
    Link,
    Snackbar,
    TextField,
    Typography,
} from "@mui/material";
import ModeEditIcon from "@mui/icons-material/ModeEdit";
import DeleteIcon from "@mui/icons-material/Delete";
import React, { useState } from "react";

type Props = {
    post: Post;
    user_id: number | null;
    topic_name: string;
    onDeletePost: () => void;
};

const useStyles = makeStyles(() => ({
    postTitle: {
        fontSize: 32,
        whiteSpace: "pre-wrap",
        paddingBottom: "0.5em",
        textAlign: "left",
        fontWeight: 500,
    },
    postBody: {
        fontSize: 16,
        textAlign: "left",
        lineHeight: 1.6,
        letterSpacing: "0.01em",
        wordBreak: "break-word",
    },
    postCard: {
        maxWidth: 800,
        width: "100%",
        marginBottom: "3em",
        padding: "8px",
        display: "flex",
        flexDirection: "column",
        minHeight: 250,
    },
    metadata: {
        fontSize: "16px !important",
        marginLeft: "auto",
    },
    cardActions: {
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
        marginTop: "auto",
    },
}));

const PostItem: React.FC<Props> = ({ post, user_id, topic_name, onDeletePost }) => {
    if (!user_id) {
        return;
    }

    const classes = useStyles();

    const [cur_post, setCurPost] = useState(post);
    const [deletePostOpen, setDeletePostOpen] = useState(false);
    const [deletePostError, setDeletePostError] = useState("");
    const [editPostOpen, setEditPostOpen] = useState(false);
    const [editPostError, setEditPostError] = useState("");
    const [showEditPostSuccess, setShowEditPostSuccess] = useState(false);

    const handleEditPostClickOpen = () => {
        setEditPostOpen(true);
    };

    const handleEditPostClose = () => {
        setEditPostOpen(false);
    };

    const handleDeletePostClickOpen = () => {
        setDeletePostOpen(true);
    };

    const handleDeletePostClose = () => {
        setDeletePostOpen(false);
    };

    const handleEditPostSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        // prevent page reload
        event.preventDefault();

        // clear previous error
        setEditPostError("");

        // get form data
        const formData = new FormData(event.currentTarget);
        const formJson = Object.fromEntries(formData.entries() as Iterable<[string, FormDataEntryValue]>);
        const post_title = formJson["post-title"];
        const post_content = formJson["post-content"];

        try {
            // Request to PATCH current post
            const response = await fetch(
                "http://localhost:8000/api/topics/" + cur_post.topic_id.toString() + "/posts/" + cur_post.id.toString(),
                {
                    method: "PATCH",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ title: post_title, content: post_content, user_id: user_id }),
                },
            );
            const data_json = await response.json();

            // process data to get result
            if (data_json.success && data_json.payload?.data) {
                const post_result = data_json.payload.data;

                if (post_result.success && post_result.post) {
                    const new_post: Post = {
                        id: post_result.post.id,
                        title: post_result.post.title,
                        user_id: post_result.post.user_id,
                        topic_id: post_result.post.topic_id,
                        content: post_result.post.content,
                        author: post_result.post.author,
                        timestamp: new Date(post_result.post.created_at),
                    };

                    // update current post, close dialog then show success message
                    setCurPost(new_post);
                    setEditPostOpen(false);
                    setShowEditPostSuccess(true);
                } else {
                    // print client-oriented error message
                    const error_message = post_result.error;
                    setEditPostError(error_message);
                }
            } else {
                console.error("Failed to PATCH post " + cur_post.id.toString() + ": %s", data_json.error);
                setEditPostError("Server error :(");
            }
        } catch (error) {
            console.error("Error fetching post " + cur_post.id.toString() + ": ", error);
            setEditPostError("Network error :(");
        }
    };

    const handleDeletePostComfirm = async () => {
        // clear previous error
        setDeletePostError("");
        try {
            // Request to DELETE current post
            const response = await fetch(
                "http://localhost:8000/api/topics/" + cur_post.topic_id.toString() + "/posts/" + cur_post.id.toString(),
                {
                    method: "DELETE",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ user_id: user_id }),
                },
            );
            const data_json = await response.json();

            // process data to get result
            if (data_json.success && data_json.payload?.data) {
                const post_result = data_json.payload.data;

                if (post_result.success) {
                    // delete current post, close dialog
                    setDeletePostOpen(false);
                    onDeletePost();
                } else {
                    // print client-oriented error message
                    const error_message = post_result.error;
                    setDeletePostError(error_message);
                }
            } else {
                console.error("Failed to Delete post " + cur_post.id.toString() + ": %s", data_json.error);
                setDeletePostError("Server error :(");
            }
        } catch (error) {
            console.error("Error fetching post " + cur_post.id.toString() + ": ", error);
            setDeletePostError("Network error :(");
        }
    };

    return (
        <>
            <Snackbar
                open={showEditPostSuccess}
                anchorOrigin={{ vertical: "top", horizontal: "center" }}
                autoHideDuration={3000}
                onClose={() => setShowEditPostSuccess(false)}
            >
                <Alert
                    onClose={() => setShowEditPostSuccess(false)}
                    severity="success"
                    variant="filled"
                    sx={{ width: "100%" }}
                >
                    Topic Modified!
                </Alert>
            </Snackbar>
            <Box sx={{ maxWidth: 1000, width: "100%", mx: "auto", px: 2 }}>
                <Box sx={{ maxWidth: 800, mx: "auto", width: "100%", mb: 1, textAlign: "left" }}>
                    <Typography color="textSecondary" fontSize={15}>
                        <Link href="/topics">Topics</Link>
                        {" > "}
                        <Link href={`/topics/${cur_post.topic_id}/posts`}>{topic_name}</Link>
                        {" > "}
                        {cur_post.title}
                    </Typography>
                </Box>
                <Box
                    sx={{
                        display: "flex",
                        flexDirection: "column",
                        alignItems: "center",
                        width: "100%",
                        mt: 1,
                    }}
                >
                    <Card className={classes.postCard}>
                        <CardContent>
                            <Box display="flex" alignItems="center" justifyContent="space-between" gap={2}>
                                <Typography variant="h5" className={classes.postTitle}>
                                    {cur_post.title}
                                </Typography>
                                {cur_post.user_id === user_id && (
                                    <Box display="flex" gap={0.5} sx={{ alignSelf: "flex-start", mt: -0.5, mr: -1 }}>
                                        <IconButton
                                            size="small"
                                            aria-label="edit topic"
                                            sx={{ p: 0.25 }}
                                            onClick={handleEditPostClickOpen}
                                        >
                                            <ModeEditIcon fontSize="small" />
                                        </IconButton>
                                        <IconButton
                                            size="small"
                                            aria-label="delete topic"
                                            sx={{ p: 0.25 }}
                                            onClick={handleDeletePostClickOpen}
                                        >
                                            <DeleteIcon fontSize="small" />
                                        </IconButton>
                                    </Box>
                                )}
                            </Box>
                            <Typography variant="body2" color="textPrimary" className={classes.postBody} component="p">
                                {cur_post.content}
                            </Typography>
                        </CardContent>
                        <CardActions className={classes.cardActions}>
                            <Box sx={{ flex: 1, display: "flex", justifyContent: "space-between", width: "100%" }}>
                                <Box sx={{ textAlign: "left" }}>
                                    <Typography color="textSecondary" className={classes.metadata}>
                                        {"Topics > " + topic_name}
                                    </Typography>
                                </Box>
                                <Box sx={{ textAlign: "right" }}>
                                    <Typography color="textSecondary" className={classes.metadata}>
                                        {"Posted by "}
                                        <strong>{cur_post.author}</strong>
                                        {" on " + cur_post.timestamp.toLocaleString()}
                                    </Typography>
                                </Box>
                            </Box>
                        </CardActions>
                    </Card>
                </Box>

                <Dialog
                    open={editPostOpen}
                    onClose={handleEditPostClose}
                    maxWidth="md"
                    fullWidth
                    sx={{ "& .MuiDialog-paper": { width: "100%", maxWidth: 700 } }}
                >
                    <DialogTitle>Edit Post</DialogTitle>
                    <DialogContent>
                        <form onSubmit={handleEditPostSubmit} id="edit-post-form" style={{ paddingTop: 2 }}>
                            <TextField
                                margin="dense"
                                id="post-title"
                                name="post-title"
                                label="Post Title"
                                defaultValue={cur_post.title}
                                fullWidth
                                variant="standard"
                            />
                            <TextField
                                margin="dense"
                                id="post-content"
                                name="post-content"
                                label="Content"
                                defaultValue={cur_post.content}
                                type="text"
                                fullWidth
                                multiline
                                InputLabelProps={{ shrink: true }}
                                minRows={8}
                                sx={{ mt: 4 }}
                            />
                        </form>
                    </DialogContent>

                    {editPostError && (
                        <Typography color="error" variant="body2" sx={{ mt: 1, textAlign: "center" }}>
                            {editPostError}
                        </Typography>
                    )}

                    <DialogActions>
                        <Button onClick={handleEditPostClose}>Cancel</Button>
                        <Button type="submit" form="edit-post-form">
                            Submit
                        </Button>
                    </DialogActions>
                </Dialog>

                <Dialog
                    open={deletePostOpen}
                    onClose={handleDeletePostClose}
                    aria-labelledby="alert-dialog-title"
                    aria-describedby="alert-dialog-description"
                >
                    <DialogTitle id="alert-dialog-title">{"Delete Post"}</DialogTitle>
                    <DialogContent>
                        <DialogContentText id="alert-dialog-description">
                            Would you like to permanently delete this post?
                        </DialogContentText>
                    </DialogContent>

                    {deletePostError && (
                        <Typography color="error" variant="body2" sx={{ mt: 1, textAlign: "center" }}>
                            {deletePostError}
                        </Typography>
                    )}

                    <DialogActions>
                        <Button onClick={handleDeletePostClose}>Back</Button>
                        <Button onClick={handleDeletePostComfirm} autoFocus>
                            Comfirm
                        </Button>
                    </DialogActions>
                </Dialog>
            </Box>
        </>
    );
};

export default PostItem;
