import Topic from "../types/Topic";
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
    Snackbar,
    TextField,
    Typography,
} from "@mui/material";
import ModeEditIcon from "@mui/icons-material/ModeEdit";
import DeleteIcon from "@mui/icons-material/Delete";
import React, { useState } from "react";
import { Link } from "react-router-dom";

type Props = {
    topic: Topic;
    user_id: number;
    onDeleteTopic: (topic_id: number) => void;
};

const API_URL = process.env.REACT_APP_API_URL;

const useStyles = makeStyles(() => ({
    topicName: {
        fontSize: 32,
        whiteSpace: "pre-wrap",
        paddingBottom: "1em",
        textAlign: "left",
    },
    topicDescription: {
        fontSize: 16,
        whiteSpace: "pre-wrap",
        paddingBottom: "1em",
        textAlign: "left",
        wordBreak: "break-word",
    },
    topicCard: {
        width: "100%",
        maxWidth: 800,
        marginBottom: "0.75rem",
        padding: "6px",
    },
    metadata: {
        fontSize: "12px !important",
        marginLeft: "auto",
    },
    cardActions: {
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
    },
}));

const TopicItem: React.FC<Props> = ({ topic, user_id, onDeleteTopic }) => {
    const classes = useStyles();
    const [cur_topic, setCurTopic] = useState(topic);
    const [editOpen, setEditOpen] = useState(false);
    const [deleteOpen, setDeleteOpen] = useState(false);
    const [editError, setEditError] = useState("");
    const [deleteError, setDeleteError] = useState("");
    const [showEditSuccess, setShowEditSuccess] = useState(false);

    const handleDeleteClickOpen = () => {
        setDeleteOpen(true);
    };

    const handleDeleteClose = () => {
        setDeleteOpen(false);
    };
    const handleEditClickOpen = () => {
        setEditOpen(true);
    };

    const handleEditClose = () => {
        setEditOpen(false);
    };

    const handleEditSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        // prevent page reload
        event.preventDefault();

        // clear previous error
        setEditError("");

        // get form data
        const formData = new FormData(event.currentTarget);
        const formJson = Object.fromEntries(formData.entries() as Iterable<[string, FormDataEntryValue]>);
        const topic_description = formJson["topic-description"];

        try {
            // Request to PATCH current topic
            const response = await fetch(`${API_URL}/api/topics/` + cur_topic.id.toString(), {
                method: "PATCH",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ description: topic_description, user_id: user_id }),
            });
            const data_json = await response.json();

            // process data to get result
            if (data_json.success && data_json.payload?.data) {
                const topic_result = data_json.payload.data;

                if (topic_result.success && topic_result.topic) {
                    const new_topic: Topic = {
                        id: topic_result.topic.id,
                        name: topic_result.topic.name,
                        user_id: topic_result.topic.user_id,
                        description: topic_result.topic.description,
                        author: topic_result.topic.author,
                        timestamp: new Date(topic_result.topic.created_at),
                    };

                    // update current topic, close dialog then show success message
                    setCurTopic(new_topic);
                    setEditOpen(false);
                    setShowEditSuccess(true);
                } else {
                    // print client-oriented error message
                    const error_message = topic_result.error;
                    setEditError(error_message);
                }
            } else {
                console.error("Failed to PATCH topic " + cur_topic.id.toString() + ": %s", data_json.error);
                setEditError("Server error :(");
            }
        } catch (error) {
            console.error("Error fetching topic " + cur_topic.id.toString() + ": ", error);
            setEditError("Network error :(");
        }
    };

    const handleDeleteComfirm = async () => {
        // clear previous error
        setDeleteError("");
        try {
            // Request to DELETE current topic
            const response = await fetch(`${API_URL}/api/topics/` + cur_topic.id.toString(), {
                method: "DELETE",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ user_id: user_id }),
            });
            const data_json = await response.json();

            // process data to get result
            if (data_json.success && data_json.payload?.data) {
                const topic_result = data_json.payload.data;

                if (topic_result.success) {
                    // delete current topic, update topic list, close dialog then show success message
                    setDeleteOpen(false);
                    onDeleteTopic(cur_topic.id);
                } else {
                    // print client-oriented error message
                    const error_message = topic_result.error;
                    setDeleteError(error_message);
                }
            } else {
                console.error("Failed to Delete topic " + cur_topic.id.toString() + ": %s", data_json.error);
                setDeleteError("Server error :(");
            }
        } catch (error) {
            console.error("Error fetching topic " + cur_topic.id.toString() + ": ", error);
            setDeleteError("Network error :(");
        }
    };

    return (
        <>
            <Snackbar
                open={showEditSuccess}
                anchorOrigin={{ vertical: "top", horizontal: "center" }}
                autoHideDuration={3000}
                onClose={() => setShowEditSuccess(false)}
            >
                <Alert
                    onClose={() => setShowEditSuccess(false)}
                    severity="success"
                    variant="filled"
                    sx={{ width: "100%" }}
                >
                    Topic Modified!
                </Alert>
            </Snackbar>

            <Card className={classes.topicCard}>
                <CardContent>
                    <Box display="flex" alignItems="center" justifyContent="space-between" gap={2}>
                        <Typography variant="h5" sx={{ pl: 1 }} className={classes.topicName}>
                            {cur_topic.name}
                        </Typography>

                        {cur_topic.user_id === user_id && (
                            <Box display="flex" gap={0.5} sx={{ alignSelf: "flex-start", mt: -0.5, mr: -1 }}>
                                <IconButton
                                    size="small"
                                    aria-label="edit topic"
                                    sx={{ p: 0.25 }}
                                    onClick={handleEditClickOpen}
                                >
                                    <ModeEditIcon fontSize="small" />
                                </IconButton>
                                <IconButton
                                    size="small"
                                    aria-label="delete topic"
                                    sx={{ p: 0.25 }}
                                    onClick={handleDeleteClickOpen}
                                >
                                    <DeleteIcon fontSize="small" />
                                </IconButton>
                            </Box>
                        )}
                    </Box>
                    <Typography
                        variant="body2"
                        color="textPrimary"
                        sx={{ px: 3 }}
                        className={classes.topicDescription}
                        component="p"
                    >
                        {cur_topic.description}
                    </Typography>
                </CardContent>
                <CardActions className={classes.cardActions}>
                    <Button size="small" sx={{ pl: 1 }} component={Link} to={`/topics/${cur_topic.id}/posts`}>
                        View Posts
                    </Button>
                    <Typography color="textSecondary" sx={{ pr: 1 }} className={classes.metadata}>
                        {"Created by "}
                        <strong>{cur_topic.author}</strong>
                        {" on " + cur_topic.timestamp.toLocaleString()}
                    </Typography>
                </CardActions>
            </Card>

            <Dialog
                open={editOpen}
                onClose={handleEditClose}
                maxWidth="md"
                fullWidth
                sx={{ "& .MuiDialog-paper": { width: "100%", maxWidth: 700 } }}
            >
                <DialogTitle>Edit Topic</DialogTitle>
                <DialogContent>
                    <form onSubmit={handleEditSubmit} id="edit-topic-form" style={{ paddingTop: 2 }}>
                        <TextField
                            disabled
                            margin="dense"
                            id="topic-name"
                            name="topic-name"
                            label="Topic Name"
                            defaultValue={cur_topic.name}
                            fullWidth
                            variant="standard"
                        />
                        <TextField
                            margin="dense"
                            id="topic-description"
                            name="topic-description"
                            label="Description"
                            defaultValue={cur_topic.description}
                            type="text"
                            fullWidth
                            multiline
                            InputLabelProps={{ shrink: true }}
                            minRows={2}
                            sx={{ mt: 4 }}
                        />
                    </form>
                </DialogContent>

                {editError && (
                    <Typography color="error" variant="body2" sx={{ mt: 1, textAlign: "center" }}>
                        {editError}
                    </Typography>
                )}

                <DialogActions>
                    <Button onClick={handleEditClose}>Cancel</Button>
                    <Button type="submit" form="edit-topic-form">
                        Submit
                    </Button>
                </DialogActions>
            </Dialog>

            <Dialog
                open={deleteOpen}
                onClose={handleDeleteClose}
                aria-labelledby="alert-dialog-title"
                aria-describedby="alert-dialog-description"
            >
                <DialogTitle id="alert-dialog-title">{"Delete Topic"}</DialogTitle>
                <DialogContent>
                    <DialogContentText id="alert-dialog-description">
                        Would you like to permanently delete this topic?
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

export default TopicItem;
