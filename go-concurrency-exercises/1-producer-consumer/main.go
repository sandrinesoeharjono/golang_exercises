//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, chan_tweets chan<- *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			// If we reach the end of the dataset, close the channel
			fmt.Println("No more tweets to produce!")
			close(chan_tweets)
			break
		}
		// Pass the individual tweets through the channel, directly to the consumer
		chan_tweets <- tweet
	}
}

func consumer(chan_tweets <-chan *Tweet) {
	// Iterate over tweets in channel of tweets
	for {
		tweet, err := <-chan_tweets
		if !err {
			fmt.Println("No more tweets to consume!")
			break
		}

		if tweet.IsTalkingAboutGo() {
			fmt.Println(tweet.Username, "\ttweets about golang")
		} else {
			fmt.Println(tweet.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	// Create channel of tweets, to stream from producer->consumer
	chan_tweets := make(chan *Tweet)

	// Producer fills channel with tweets, as it goes
	go producer(stream, chan_tweets)

	// Consumer processes tweets as they come through the channel
	consumer(chan_tweets)

	fmt.Printf("Process took %s\n", time.Since(start))
}
