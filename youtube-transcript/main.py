import logging
from utils import retrieve_video_id, merge_transcript_into_full_text

ACCEPTED_LANGUAGE = ["en", "en-US", "en-UK"]
# maybe have workaround for this

"""
Checks whether the given subtitle language is English.

Args:
    sub_lang (str): The subtitle language to check.

Returns:
    bool: True if the subtitle language is English, False otherwise.
"""
def is_eng_sub(sub_lang: str) -> bool:
    candidates = ["en"]
    for prefix in candidates:
        if sub_lang.startswith(prefix):
            return True
    return False


from youtube_transcript_api import YouTubeTranscriptApi
def get_transcript(youtube_url: str) -> str:
    id = retrieve_video_id(youtube_url)
    
    transcript_list = YouTubeTranscriptApi.list_transcripts(id)
    
    transcript = transcript_list.find_manually_created_transcript(ACCEPTED_LANGUAGE)

    full = merge_transcript_into_full_text(transcript)
    
    return full


        
if __name__ == "__main__":
    URL = "https://www.youtube.com/watch?v=QdnVT22LBBs"

    trans = get_transcript(URL) 
    print(trans)