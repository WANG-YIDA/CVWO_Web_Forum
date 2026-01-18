import Comment from "../types/Comment";
import DeleteIcon from "@mui/icons-material/Delete";
import {
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
    Typography,
} from "@mui/material";
import { makeStyles } from "@mui/styles";
import React, { useState } from "react";

type Props = {
    comment: Comment;
    user_id: number | null;
    topic_id: string | null;
    onDeleteComment: (comment_id: number) => void;
};
const useStyles = makeStyles(() => ({
    commentBody: {
        fontSize: 16,
        whiteSpace: "pre-wrap",
        paddingBottom: "0.4em",
        textAlign: "left",
        wordBreak: "break-word",
    },
    commentCard: {
        width: "100%",
        maxWidth: 750,
        marginBottom: "0.75rem",
        padding: "3px",
        marginLeft: 0,
        marginRight: 0,
    },
    cardContent: {
        display: "flex",
        flexDirection: "column",
        alignItems: "flex-start",
        gap: "0.5rem",
    },
    metadata: {
        fontSize: "12px !important",
        alignSelf: "flex-end",
        marginTop: "auto",
    },
    cardActions: {
        display: "flex",
        justifyContent: "flex-end",
        alignItems: "center",
        padding: "0 16px 3px",
    },
}));

const CommentItem: React.FC<Props> = ({ comment, user_id, topic_id, onDeleteComment }) => {
    if (!user_id) {
        return;
    }

    const classes = useStyles();
    const [deleteOpen, setDeleteOpen] = useState(false);
    const [deleteError, setDeleteError] = useState("");

    const handleDeleteClickOpen = () => {
        setDeleteOpen(true);
    };

    const handleDeleteClose = () => {
        setDeleteOpen(false);
    };

    const handleDeleteComfirm = async () => {
        // clear previous error message
        setDeleteError("");

        try {
            // Request to DELETE current comment
            const response = await fetch(
                "http://localhost:8000/api/topics/" +
                    topic_id +
                    "/posts/" +
                    comment.post_id.toString() +
                    "/comments/" +
                    comment.id.toString(),
                {
                    method: "DELETE",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ user_id: user_id }),
                },
            );
            const data_json = await response.json();

            // process data to get result
            if (data_json.success && data_json.payload?.data) {
                const comment_result = data_json.payload.data;

                if (comment_result.success) {
                    // delete current comment, update comment list, close dialog then show success message
                    setDeleteOpen(false);
                    onDeleteComment(comment.id);
                } else {
                    // print client-oriented error message
                    const error_message = comment_result.error;
                    setDeleteError(error_message);
                }
            } else {
                console.error("Failed to Delete comment " + comment.id.toString() + ": %s", data_json.error);
                setDeleteError("Server error :(");
            }
        } catch (error) {
            console.error("Error fetching comment " + comment.id.toString() + ": ", error);
            setDeleteError("Network error :(");
        }
    };

    return (
        <>
            <Card className={classes.commentCard}>
                <CardContent className={classes.cardContent}>
                    <Box
                        sx={{
                            width: "100%",
                            maxWidth: "100%",
                            padding: "0 16px",
                            margin: "0 auto",
                            position: "relative",
                        }}
                    >
                        {user_id === comment.user_id && (
                            <IconButton
                                size="small"
                                aria-label="delete comment"
                                sx={{ p: 0.25, position: "absolute", top: -5, right: 25 }}
                                onClick={handleDeleteClickOpen}
                            >
                                <DeleteIcon fontSize="small" />
                            </IconButton>
                        )}

                        <Typography
                            variant="body2"
                            color="textPrimary"
                            sx={{ mt: 1, pr: 6 }}
                            className={classes.commentBody}
                            component="p"
                        >
                            {comment.content}
                        </Typography>
                    </Box>
                </CardContent>
                <CardActions className={classes.cardActions}>
                    <Typography color="textSecondary" className={classes.metadata} gutterBottom sx={{ mt: -0.25 }}>
                        {"Commented by "}
                        <strong>{comment.author}</strong>
                        {" on " + comment.timestamp.toLocaleString()}
                    </Typography>
                </CardActions>
            </Card>

            <Dialog
                open={deleteOpen}
                onClose={handleDeleteClose}
                aria-labelledby="alert-dialog-title"
                aria-describedby="alert-dialog-description"
            >
                <DialogTitle id="alert-dialog-title">{"Delete Comment"}</DialogTitle>
                <DialogContent>
                    <DialogContentText id="alert-dialog-description">
                        Would you like to permanently delete this comment?
                    </DialogContentText>
                </DialogContent>

                {deleteError && (
                    <Typography color="error" variant="body2" sx={{ mt: 1, textAlign: "center" }}>
                        {deleteError}
                    </Typography>
                )}

                <DialogActions>
                    <Button onClick={handleDeleteClose}>Back</Button>
                    <Button onClick={handleDeleteComfirm} autoFocus>
                        Comfirm
                    </Button>
                </DialogActions>
            </Dialog>
        </>
    );
};

export default CommentItem;
