import TopicList from "../components/TopicList";
import Topic from "../types/Topic";
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
    Divider,
    IconButton,
    keyframes,
    Snackbar,
    TextField,
    Typography,
} from "@mui/material";
import React, { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";

interface TopicJSON {
    id: number;
    name: string;
    description: string;
    user_id: number;
    author: string;
    created_at: string;
}

const API_DOMAIN = process.env.REACT_APP_API_DOMAIN;

const fadeIn = keyframes`
    from { opacity: 0; transform: translateY(12px); }
    to { opacity: 1; transform: translateY(0); }
`;

const TopicListView: React.FC = () => {
    const [topics, setTopics] = useState<Topic[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [open, setOpen] = useState(false);
    const [error, setError] = useState("");
    const [userID, setUserID] = useState<number | null>(null);
    const [showSuccess, setShowSuccess] = useState(false);
    const [showDeleteSuccess, setShowDeleteSuccess] = useState(false);
    const [server_error, setServerError] = useState(false);
    const navigate = useNavigate();

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
        const topic_name = formJson["topic-name"];
        const topic_description = formJson["topic-description"];

        try {
            // Request to POST topics
            const response = await fetch(`${API_DOMAIN}:/api/topics`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ name: topic_name, description: topic_description, user_id: userID }),
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

                    // put new topic at top, close dialog then show success message
                    setTopics((topics) => [new_topic, ...topics]);
                    setOpen(false);
                    setShowSuccess(true);
                } else {
                    // print client-oriented error message
                    const error_message = topic_result.error;
                    setError(error_message);
                }
            } else {
                console.error("Failed to POST topics: %s", data_json.error);
                setError("Server error :(");
            }
        } catch (error) {
            console.error("Error fetching topics:", error);
            setError("Network error :(");
        }
    };

    const handleDeleteTopic = (topic_id: number) => {
        setTopics((topics) => topics.filter((topic) => topic.id !== topic_id));
        setShowDeleteSuccess(true);
    };

    // navigate to login page if not login yet
    useEffect(() => {
        const user_id = sessionStorage.getItem("user_id");
        if (!user_id) {
            navigate("/login");
        } else {
            setUserID(parseInt(user_id, 10));
        }
    }, [navigate]);

    useEffect(() => {
        const fetchTopics = async () => {
            try {
                // Request to GET Topics
                const response = await fetch(`${API_DOMAIN}:/api/topics`, {
                    method: "GET",
                    headers: { "Content-Type": "application/json" },
                });
                const data_json = await response.json();

                // process data to get topics
                if (data_json.success && data_json.payload?.data) {
                    const topicListResult = data_json.payload.data;

                    if (topicListResult.success && topicListResult.topics) {
                        const topicList: Topic[] = topicListResult.topics.map((topic: TopicJSON) => ({
                            id: topic.id,
                            name: topic.name,
                            description: topic.description,
                            user_id: topic.user_id,
                            author: topic.author,
                            timestamp: new Date(topic.created_at),
                        }));

                        setTopics(topicList.reverse());
                    }
                } else {
                    console.error("Failed to GET topics: %s", data_json.error);
                    setServerError(true);
                }
            } catch (error) {
                console.error("Error fetching topics:", error);
                setServerError(true);
            } finally {
                setLoading(false);
            }
        };

        fetchTopics();
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
                alignItems: "center",
                minHeight: "100vh",
                textAlign: "center",
                gap: "3rem",
            }}
        >
            <Box sx={{ width: "100%", maxWidth: 800, padding: "0 16px", margin: "0 auto" }}>
                <Typography
                    component="h4"
                    sx={{
                        whiteSpace: "nowrap",
                        fontSize: { xs: "2.5rem", sm: "3rem", md: "3.5rem" },
                        paddingTop: "4.5rem",
                        mb: "3rem",
                        background: "linear-gradient(45deg, #2E3192, #1BFFFF)",
                        WebkitBackgroundClip: "text",
                        WebkitTextFillColor: "transparent",
                        backgroundClip: "text",
                        filter: "drop-shadow(2px 2px 4px rgba(0, 0, 0, 0.3))",
                        opacity: 0,
                        animation: `${fadeIn} 1.2s ease forwards`,
                    }}
                >
                    <strong>Discover Your Conversations</strong>
                </Typography>

                {loading ? (
                    <Typography sx={{ mt: 4, color: "text.secondary" }}>Loading topics...</Typography>
                ) : (
                    <>
                        <Box sx={{ display: "flex", justifyContent: "space-between", mr: 1 }}>
                            <Button component={Link} to="/">
                                {"< Home"}
                            </Button>
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
                                New Topic Added!
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
                                Topic Deleted!
                            </Alert>
                        </Snackbar>

                        <Divider sx={{ borderColor: "#bababeff", borderBottomWidth: 2 }} />

                        <TopicList topics={topics} user_id={userID} onDeleteTopic={handleDeleteTopic} />
                    </>
                )}
            </Box>

            <Dialog
                open={open}
                onClose={handleClose}
                maxWidth="md"
                fullWidth
                sx={{ "& .MuiDialog-paper": { width: "100%", maxWidth: 700 } }}
            >
                <DialogTitle>Create New Topic</DialogTitle>
                <DialogContent>
                    <DialogContentText>Start a new conversation and invite others to join!</DialogContentText>
                    <form onSubmit={handleSubmit} id="create-topic-form" style={{ paddingTop: 2 }}>
                        <TextField
                            autoFocus
                            required
                            margin="dense"
                            id="topic-name"
                            name="topic-name"
                            label="Topic Name"
                            type="text"
                            fullWidth
                            variant="standard"
                        />
                        <TextField
                            margin="dense"
                            id="topic-description"
                            name="topic-description"
                            label="Description"
                            type="text"
                            fullWidth
                            multiline
                            InputLabelProps={{ shrink: true }}
                            minRows={2}
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
                    <Button type="submit" form="create-topic-form">
                        Create
                    </Button>
                </DialogActions>
            </Dialog>
        </div>
    );
};

export default TopicListView;
