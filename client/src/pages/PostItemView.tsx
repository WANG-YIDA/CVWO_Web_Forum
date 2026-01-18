import Post from "../types/Post";
import Comment from "../types/Comment";
import CommentList from "../components/CommentList";
import PostItem from "../components/PostItem";
import React, { useEffect, useState } from "react";
import { Alert, Box, Button, Snackbar, TextField, Typography } from "@mui/material";
import { useNavigate, useParams } from "react-router-dom";

interface CommentJSON {
    id: number;
    content: string;
    user_id: number;
    post_id: number;
    author: string;
    created_at: string;
}

const PostItemView: React.FC = () => {
    const [topicName, setTopicName] = useState("");
    const [post, setPost] = useState<Post>();
    const [comments, setComments] = useState<Comment[]>([]);
    const [commentInput, setCommentInput] = useState("");
    const [showCreateCommentSuccess, setShowCreateCommentSuccess] = useState(false);
    const [showDeleteCommentSuccess, setShowDeleteCommentSuccess] = useState(false);
    const [createCommentError, setCreateCommentError] = useState("");
    const [userID, setUserID] = useState<number | null>(null);
    const [server_error, setServerError] = useState(false);
    const { topicID } = useParams<{ topicID: string }>(); // Get topicID from URL params
    const { postID } = useParams<{ postID: string }>(); // Get postID from URL params
    const navigate = useNavigate();

    const handleCreateCommentSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        // prevent page reload
        event.preventDefault();

        // clear previous error
        setCreateCommentError("");

        // get form data
        const formData = new FormData(event.currentTarget);
        const formJson = Object.fromEntries(formData.entries() as Iterable<[string, FormDataEntryValue]>);
        const comment_content = formJson["comment-content"];

        try {
            // Request to POST comments
            const response = await fetch(
                "http://localhost:8000/api/topics/" + topicID + "/posts/" + postID + "/comments",
                {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ content: comment_content, user_id: userID }),
                },
            );
            const data_json = await response.json();

            // process data to get result
            if (data_json.success && data_json.payload?.data) {
                const comment_result = data_json.payload.data;

                if (comment_result.success && comment_result.comment) {
                    const new_comment: Comment = {
                        id: comment_result.comment.id,
                        post_id: comment_result.comment.post_id,
                        user_id: comment_result.comment.user_id,
                        content: comment_result.comment.content,
                        author: comment_result.comment.author,
                        timestamp: new Date(comment_result.comment.created_at),
                    };

                    // put new comment at top, close dialog then show success message
                    setCommentInput("");
                    setComments((comments) => [new_comment, ...comments]);
                    setShowCreateCommentSuccess(true);
                } else {
                    // print client-oriented error message
                    const error_message = comment_result.error;
                    setCreateCommentError(error_message);
                }
            } else {
                console.error("Failed to POST comments: %s", data_json.error);
                setCreateCommentError("Server error :(");
            }
        } catch (error) {
            console.error("Error fetching comments:", error);
            setCreateCommentError("Network error :(");
        }
    };

    // navigate to login page if not login yet, or to topics/posts page if topic/post id is not valid
    useEffect(() => {
        const user_id = sessionStorage.getItem("user_id");
        if (!user_id) {
            navigate("/login");
        } else if (!topicID) {
            navigate("/topics");
        } else if (!postID) {
            navigate(`/topics/${topicID}/posts`);
        } else {
            setUserID(parseInt(user_id, 10));
        }
    }, [navigate]);

    // Get Post
    useEffect(() => {
        const fetchPost = async () => {
            try {
                // Request to GET post
                const response = await fetch("http://localhost:8000/api/topics/" + topicID + "/posts/" + postID, {
                    method: "GET",
                    headers: { "Content-Type": "application/json" },
                });
                const data_json = await response.json();

                // process data to get post
                if (data_json.success && data_json.payload?.data) {
                    const postResult = data_json.payload.data;

                    if (postResult.success && postResult.post) {
                        const post: Post = {
                            id: postResult.post.id,
                            title: postResult.post.title,
                            topic_id: postResult.post.topic_id,
                            content: postResult.post.content,
                            user_id: postResult.post.user_id,
                            author: postResult.post.author,
                            timestamp: new Date(postResult.post.created_at),
                        };

                        setPost(post);
                    }
                } else {
                    console.error("Failed to GET post: %s", data_json.error);
                    setServerError(true);
                }
            } catch (error) {
                console.error("Error fetching post:", error);
                setServerError(true);
            }
        };

        const fetchComments = async () => {
            try {
                // Request to GET comments
                const response = await fetch(
                    "http://localhost:8000/api/topics/" + topicID + "/posts/" + postID + "/comments",
                    {
                        method: "GET",
                        headers: { "Content-Type": "application/json" },
                    },
                );
                const data_json = await response.json();

                // process data to get comments
                if (data_json.success && data_json.payload?.data) {
                    const commentListResult = data_json.payload.data;

                    if (commentListResult.success && commentListResult.comments) {
                        const commentList: Comment[] = commentListResult.comments.map((comment: CommentJSON) => ({
                            id: comment.id,
                            post_id: comment.post_id,
                            content: comment.content,
                            user_id: comment.user_id,
                            author: comment.author,
                            timestamp: new Date(comment.created_at),
                        }));

                        setComments(commentList);
                    }
                } else {
                    console.error("Failed to GET post: %s", data_json.error);
                    setServerError(true);
                }
            } catch (error) {
                console.error("Error fetching post:", error);
                setServerError(true);
            }
        };

        // get topic name
        const fetchTopic = async () => {
            try {
                // Request to GET topic
                const response = await fetch("http://localhost:8000/api/topics/" + topicID, {
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
                    console.error("Failed to GET topic: %s", data_json.error);
                    setServerError(true);
                }
            } catch (error) {
                console.error("Error fetching topic:", error);
                setServerError(true);
            }
        };

        fetchTopic();
        fetchPost();
        fetchComments();
    }, []);

    const handleDeleteComment = (comment_id: number) => {
        setComments((comments) => comments.filter((comment) => comment.id !== comment_id));
        setShowDeleteCommentSuccess(true);
    };

    const handleDeletePost = () => {
        navigate("/topics/" + topicID + "/posts");
    };

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
                alignItems: "center",
                minHeight: "100vh",
                textAlign: "center",
                gap: "1rem",
                paddingTop: "6rem",
            }}
        >
            {post ? (
                <PostItem post={post} user_id={userID} topic_name={topicName} onDeletePost={handleDeletePost} />
            ) : (
                <div>Loading post...</div>
            )}

            <Box sx={{ maxWidth: 800, width: "100%", mx: "auto", px: 2, mt: 4, textAlign: "left" }}>
                <Typography
                    variant="h5"
                    color="textSecondary"
                    gutterBottom
                    align="left"
                    marginBottom={2}
                    marginLeft={-1}
                >
                    <strong>Comments:</strong>
                </Typography>

                <form onSubmit={handleCreateCommentSubmit} id="create-comment-form" style={{ paddingTop: 2 }}>
                    <Box sx={{ display: "flex", alignItems: "flex-start", gap: 1 }}>
                        <TextField
                            margin="dense"
                            id="comment-content"
                            name="comment-content"
                            label="Share your thoughts here..."
                            type="text"
                            fullWidth
                            multiline
                            minRows={1}
                            sx={{
                                my: 2,
                                borderRadius: 1,
                                "& .MuiInputBase-root": {
                                    color: "#222",
                                },
                                "& .MuiOutlinedInput-notchedOutline": {
                                    borderColor: "#6097e4ff",
                                },
                                "&:hover .MuiOutlinedInput-notchedOutline": {
                                    borderColor: "#1966beff",
                                },
                                "& .MuiOutlinedInput-root.Mui-focused .MuiOutlinedInput-notchedOutline": {
                                    borderColor: "#66a3e9ff",
                                },
                                "& .MuiInputLabel-root": {
                                    color: "#1976d2",
                                },
                            }}
                            value={commentInput}
                            onChange={(e) => setCommentInput(e.target.value)}
                        />
                        <Button
                            type="submit"
                            variant="outlined"
                            sx={{
                                height: "45px",
                                mt: 2.25,
                                ml: 2,
                                mr: -2,
                                whiteSpace: "nowrap",
                                color: "#407cd1ff",
                                letterSpacing: 1,
                                borderColor: "#6097e4ff",
                                "&:hover": {
                                    backgroundColor: "#66a3e9ff",
                                    color: "#fff",
                                },
                            }}
                        >
                            Create
                        </Button>
                    </Box>
                </form>

                {createCommentError && (
                    <Typography color="error" variant="body2" sx={{ mt: 1, textAlign: "center" }}>
                        {createCommentError}
                    </Typography>
                )}
            </Box>

            <Snackbar
                open={showCreateCommentSuccess}
                anchorOrigin={{ vertical: "top", horizontal: "center" }}
                autoHideDuration={3000}
                onClose={() => setShowCreateCommentSuccess(false)}
            >
                <Alert
                    onClose={() => setShowCreateCommentSuccess(false)}
                    severity="success"
                    variant="filled"
                    sx={{ width: "100%" }}
                >
                    New Comment Added!
                </Alert>
            </Snackbar>

            <Snackbar
                open={showDeleteCommentSuccess}
                anchorOrigin={{ vertical: "top", horizontal: "center" }}
                autoHideDuration={3000}
                onClose={() => setShowDeleteCommentSuccess(false)}
            >
                <Alert
                    onClose={() => setShowDeleteCommentSuccess(false)}
                    severity="success"
                    variant="filled"
                    sx={{ width: "100%" }}
                >
                    Comment Deleted!
                </Alert>
            </Snackbar>

            <CommentList
                comments={comments}
                user_id={userID}
                topic_id={topicID || null}
                onDeleteComment={handleDeleteComment}
            />
        </div>
    );
};

export default PostItemView;
