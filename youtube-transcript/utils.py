import logging
"""
Retrieves the video ID from a given YouTube URL.

Args:
    youtube_url (str): The URL of the YouTube video.

Returns:
    str: The video ID extracted from the URL.

Raises:
    Exception: If the URL cannot be parsed to extract the video ID.
"""
from urllib.parse import urlparse, parse_qs
def retrieve_video_id(youtube_url: str) -> str:
    try:
        parse_url = urlparse(youtube_url)
        query_params = parse_qs(parse_url.query) # query_params is a dict
        if 'v' not in query_params:
            logging.info(f"{youtube_url} doesn't have id")
            return ""
        return query_params["v"][0]
    except Exception as e:
        logging.info(f"{youtube_url} can not be parsed\nexception: {e}")

"""
Merges the text from all segments in a YouTube transcript into a single string.

Args:
    transcript (YouTubeTranscriptApi.Transcript): The transcript object containing the segments.

Returns:
    str: The full text of the transcript as a single string.
"""
def merge_transcript_into_full_text(transcript) -> str:
    subs = [segment['text'] for segment in transcript.fetch()]
    return " ".join(subs)