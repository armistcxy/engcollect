import nltk
from nltk.stem import WordNetLemmatizer
nltk.download('wordnet')
lemmatizer = WordNetLemmatizer()

"""
    Lemmatizes the given list of tokens using the WordNetLemmatizer.
    
    Args:
        tokens (list[str]): The list of tokens to lemmatize.
    
    Returns:
        list[str]: The list of lemmatized tokens.
    """
def lemmatize(tokens: list[str]) -> list[str]: 
    return [lemmatizer.lemmatize(token) for token in tokens]

"""
    Removes duplicate tokens from the given list.
    
    Args:
        tokens (list[str]): The list of tokens to remove duplicates from.
    
    Returns:
        list[str]: The list of tokens with duplicates removed.
"""
def remove_duplicate(tokens: list[str]) -> list[str]: 
    return list(set(tokens))


nltk.download('stopwords')
from nltk.corpus import stopwords
stopwords = set(stopwords.words('english'))

"""
    Removes stopwords from the given list of tokens.
    
    Args:
        tokens (list[str]): The list of tokens to remove stopwords from.
    
    Returns:
        list[str]: The list of tokens with stopwords removed.
    """
def remove_stopwords(tokens: list[str]) -> list[str]: 
    return [token for token in tokens if token not in stopwords]


import re
from nltk import tokenize


"""
        Cleans the given text by performing the following steps:
        1. Converts the text to lowercase.
        2. Removes URLs from the text.
        3. Removes HTML tags from the text.
        4. Removes non-alphanumeric characters from the text.
        5. Tokenizes the text into a list of tokens.
        6. Removes duplicate tokens from the list.
        7. Removes stopwords from the list of tokens.
        8. Lemmatizes the tokens in the list.
    
        Args:
            text (str): The text to be cleaned.
    
        Returns:
            list[str]: The list of cleaned and processed tokens.
    """
def nlp_clean_pipe(text: str) -> list[str]: 
    # this may cost a significant amount of time, given that Cambridge's dictionary is case insensitive, we can remove this step
    text = text.lower() 
    
    # remove URLs
    text = re.sub(r'http\S+|www\S+|https\S+', '', text, flags=re.MULTILINE)
    
    # Remove HTML tags
    text = re.sub(r'<.*?>', '', text)
    
    # Remove non-alphanumeric characters
    text = re.sub(r'[^a-zA-Z0-9\s]', '', text)
    
    tokens = tokenize(text)
    
    # remove duplicate
    tokens = remove_duplicate(tokens)
    
    # remove stopwords
    tokens = remove_stopwords(tokens)
    
    # lemmatize
    tokens = lemmatize(tokens) 
    
    return tokens