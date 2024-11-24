import React, { useState } from "react";
import "./ContentGen.css";

// Define the type for a post
interface Post {
  platform: string;
  content: string;
}

const ContentGen: React.FC = () => {
  const [posts, setPosts] = useState<Post[]>([]); // State to store parsed posts
  const [isLoading, setIsLoading] = useState(false);

  // Helper function to clean and parse JSON strings
  const processJsonString = (jsonString: string): Post[] => {
    // Remove ```json and ``` from the string
    const cleanedString = jsonString
      .replace(/```json/g, "")
      .replace(/```/g, "")
      .trim();

    // Parse the cleaned JSON string into a JavaScript array of objects
    let parsedObjects: Post[] = [];
    try {
      parsedObjects = JSON.parse(cleanedString) as Post[];
    } catch (error) {
      console.error("Error parsing JSON string:", error);
    }

    return parsedObjects;
  };

  const generatePost = async () => {
    setIsLoading(true);
    try {
      const response = await fetch(`http://localhost:8000/generate-posts`);
      const data = await response.json();

      // Assume `data.response` is a JSON string with backticks
      const parsedPosts = processJsonString(data.data);

      setPosts(parsedPosts);
      // alert("Posts generated successfully!");
    } catch (err) {
      console.error("Error fetching or parsing the data:", err);
    } finally {
      setIsLoading(false);
    }
  };

  const handlePost = async (post: string) => {
    try {
      const response = await fetch(`http://127.0.0.1:3000/tweets`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ content: post }),
      });
      const data = await response.json();
      alert("Post successful!");
      console.log(data);
      alert("Post successful!");
    } catch (error) {
      console.error("Error posting the data:", error);
      alert("Something went wrong");
    }
  };
  return (
    <div className="Concontainer">
      <header className="Conheader">Post Generator</header>
      <p className="Condescription">
        Welcome to the Post Generator! Click the button below to create a new
        post.
      </p>
      <button className="Conbutton" onClick={generatePost} disabled={isLoading}>
        {isLoading ? "Loading..." : "Generate Post"}
      </button>
      <div className="Conposts">
        {posts.length > 0 && (
          <ul>
            {posts.map((post, index) => (
              <li key={index}>
                <strong>{post.platform}:</strong> {post.content}
                <button
                  onClick={() => {
                    handlePost(post.content);
                  }}
                >
                  post
                </button>
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
};

export default ContentGen;
