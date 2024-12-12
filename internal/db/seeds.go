package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/AhmedRabea0302/go-social/internal/store"
)

var usernames = []string{
	"Alice", "Bob", "Charlie", "Diana", "Ethan",
	"Fiona", "George", "Hannah", "Ian", "Jack",
	"Karen", "Liam", "Mia", "Nathan", "Olivia",
	"Paul", "Quinn", "Rachel", "Sam", "Tina",
	"Uma", "Victor", "Wendy", "Xander", "Yara",
	"Zane", "Abigail", "Benjamin", "Catherine", "Daniel",
	"Emma", "Frank", "Grace", "Henry", "Isabella",
	"Jacob", "Kaitlyn", "Logan", "Mason", "Nora",
	"Owen", "Penelope", "Quentin", "Rebecca", "Sophia",
	"Tyler", "Ursula", "Violet", "William", "Zoe",
}

var titles = []string{
	"10 Tips to Boost Your Productivity",
	"Understanding Cloud Computing",
	"The Beginner’s Guide to Investing",
	"How to Build a Simple Web App",
	"Top 5 Coding Practices for Developers",
	"The Future of Artificial Intelligence",
	"How to Start a Successful Blog",
	"The Benefits of Regular Exercise",
	"Traveling on a Budget: Top Destinations",
	"Mastering the Art of Public Speaking",
	"The Ultimate Guide to SEO",
	"How to Create Stunning Visual Content",
	"Understanding Cryptocurrency and Blockchain",
	"Work-Life Balance: Tips and Tricks",
	"Top 10 Books Every Entrepreneur Should Read",
	"The Rise of Remote Work: What You Need to Know",
	"How to Build a Personal Brand",
	"The Science Behind Healthy Eating",
	"Design Trends to Watch in 2024",
	"How to Learn a New Language Fast",
}

var contents = []string{
	"Discover practical tips to enhance your productivity and achieve your goals efficiently.",
	"Learn the fundamentals of cloud computing and its impact on modern technology.",
	"A beginner-friendly guide to understanding the world of investments and financial growth.",
	"Step-by-step instructions to build your first simple web application using Go.",
	"Explore the best coding practices that will elevate your skills as a developer.",
	"Dive into the advancements and possibilities in the field of artificial intelligence.",
	"Everything you need to know to launch a successful blog and attract an audience.",
	"Understand the benefits of regular exercise for your body and mind.",
	"Top destinations and tips for traveling on a budget without compromising experience.",
	"Master the essential skills of public speaking to captivate and inspire your audience.",
	"A comprehensive guide to optimizing your website for better search engine rankings.",
	"How to create visually appealing content that grabs attention and boosts engagement.",
	"An introduction to cryptocurrency and blockchain technology for beginners.",
	"Practical advice to maintain a healthy balance between your work and personal life.",
	"A curated list of must-read books for aspiring and seasoned entrepreneurs.",
	"Understand the growing trend of remote work and how it is shaping the future workplace.",
	"Strategies to build a powerful personal brand that stands out in your industry.",
	"The science-backed principles of healthy eating for a better quality of life.",
	"Stay ahead of the curve with the latest design trends expected to dominate in 2024.",
	"Learn techniques to quickly acquire a new language and become conversational.",
}

var tags = []string{
	"technology", "health", "travel", "finance", "lifestyle",
	"coding", "education", "business", "food", "productivity",
	"fitness", "design", "marketing", "art", "music",
	"photography", "DIY", "parenting", "sports", "news",
}

var comments = []string{
	"Great post! Very informative.",
	"I totally agree with your points here.",
	"Could you elaborate more on this topic?",
	"This was really helpful, thank you!",
	"I had a similar experience, thanks for sharing.",
	"This is exactly what I was looking for!",
	"Not sure I agree with this perspective.",
	"Do you have any resources to recommend?",
	"Awesome read, keep up the great work!",
	"This raises some interesting questions.",
	"Thanks for breaking this down so clearly.",
	"I learned a lot from this post!",
	"What inspired you to write this?",
	"This is such a unique take on the subject.",
	"I think this could apply to other scenarios too.",
	"Can you explain this part a bit further?",
	"This was a fun read, thanks for posting.",
	"Your insights are always valuable.",
	"I didn’t quite follow this section.",
	"Looking forward to your next post!",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	// Create some initial users
	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", user)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", post)
			return
		}

		comments := generateComments(500, users, posts)
		for _, comment := range comments {
			if err := store.Comments.Create(ctx, comment); err != nil {
				log.Println("Error creating comment:", comment)
				return
			}
		}
	}

	log.Println("Seeding Completed Successfully...")

}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)
	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123123",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		post := posts[rand.Intn(len(posts))]

		cms[i] = &store.Comment{
			UserID:  user.ID,
			PostID:  post.ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
