// Package readable provides an easy way to generate readable random phrases.
package readable

import (
	"math/rand"
	"strings"
	"time"
)

var (
	adjectives []string
	nouns      []string
	adjCount   int
	nounCount  int
	vowels     []string
)

func init() {
	adjectives = []string{"aback", "abaft", "abandoned", "abashed", "aberrant", "abhorrent", "abiding", "abject", "ablaze", "able", "abnormal", "aboard", "aboriginal", "abortive", "abounding", "abrasive", "abrupt", "absent", "absorbed", "absorbing", "abstracted", "absurd", "abundant", "abusive", "acceptable", "accessible", "accidental", "accurate", "acid", "acidic", "acoustic", "acrid", "actually", "ad", "ad", "hoc", "adamant", "adaptable", "addicted", "adhesive", "adjoining", "adorable", "adventurous", "afraid", "aggressive", "agonizing", "agreeable", "ahead", "ajar", "alcoholic", "alert", "alike", "alive", "alleged", "alluring", "aloof", "amazing", "ambiguous", "ambitious", "amuck", "amused", "amusing", "ancient", "angry", "animated", "annoyed", "annoying", "anxious", "apathetic", "aquatic", "aromatic", "arrogant", "ashamed", "aspiring", "assorted", "astonishing", "attractive", "auspicious", "automatic", "available", "average", "awake", "aware", "awesome", "awful", "axiomatic", "bad", "barbarous", "bashful", "bawdy", "beautiful", "belligerent", "berserk", "better", "bewildered", "big", "billowy", "bitter", "bizarre", "black", "bloody", "blue", "blushing", "boiling", "boorish", "bored", "boring", "boundless", "brainy", "brash", "brave", "brawny", "breakable", "breeze", "breezy", "brief", "bright", "broad", "broken", "bumpy", "burly", "busy", "cagey", "callous", "calm", "capable", "capricious", "careful", "cautious", "ceaseless", "changeable", "charming", "cheerful", "childlike", "chilly", "chivalrous", "chubby", "chunky", "clammy", "classy", "clean", "clear", "clever", "cloistered", "cloudy", "clumsy", "coherent", "cold", "colorful", "colossal", "combative", "comfortable", "concerned", "condemned", "confused", "cooing", "cool", "cooperative", "courageous", "cowardly", "crabby", "craven", "crazy", "credible", "creepy", "crooked", "crowded", "cruel", "cuddly", "cultured", "curious", "curly", "curved", "cute", "cynical", "daffy", "daily", "damaged", "damaging", "damp", "dangerous", "dapper", "dark", "dashing", "dazzling", "dead", "deadpan", "deafening", "debonair", "decisive", "decorous", "deep", "deeply", "defeated", "defective", "defiant", "delicious", "delightful", "demonic", "depressed", "deranged", "deserted", "detailed", "determined", "devilish", "didactic", "different", "difficult", "diligent", "direful", "dirty", "disagreeable", "discreet", "disgusted", "disillusioned", "dispensable", "distinct", "disturbed", "divergent", "dizzy", "domineering", "doubtful", "drab", "draconian", "dramatic", "drunk", "dry", "dull", "dusty", "dynamic", "dysfunctional", "eager", "early", "earsplitting", "earthy", "easy", "eatable", "economic", "educated", "efficacious", "efficient", "elated", "elderly", "elegant", "elfin", "elite", "embarrassed", "eminent", "empty", "enchanting", "encouraging", "endurable", "energetic", "entertaining", "enthusiastic", "envious", "equable", "erect", "erratic", "ethereal", "evanescent", "evasive", "evil", "excellent", "excited", "exclusive", "exotic", "expensive", "exuberant", "exultant", "fabulous", "faded", "faint", "fair", "faithful", "fallacious", "famous", "fanatical", "fancy", "fantastic", "fascinated", "fast", "fat", "faulty", "fearless", "feigned", "fertile", "festive", "few", "fierce", "filthy", "fine", "finicky", "flagrant", "flaky", "flashy", "flat", "flawless", "flippant", "flowery", "fluffy", "fluttering", "foamy", "foolish", "foregoing", "forgetful", "fortunate", "fragile", "frail", "frantic", "freezing", "fresh", "fretful", "friendly", "frightened", "full", "functional", "funny", "furtive", "futuristic", "fuzzy", "gabby", "gainful", "gamy", "gaping", "garrulous", "gaudy", "gentle", "giant", "giddy", "gifted", "gigantic", "glamorous", "gleaming", "glib", "glorious", "glossy", "godly", "good", "goofy", "gorgeous", "graceful", "grandiose", "gray", "greasy", "great", "greedy", "green", "grieving", "groovy", "grotesque", "grouchy", "grubby", "gruesome", "grumpy", "guarded", "guiltless", "gullible", "gusty", "guttural", "habitual", "half", "hallowed", "halting", "handsome", "handsomely", "hapless", "happy", "hard", "harmonious", "harsh", "heady", "healthy", "heartbreaking", "heavenly", "heavy", "hellish", "helpful", "helpless", "hesitant", "high", "highfalutin", "hilarious", "hissing", "historical", "hoc", "holistic", "hollow", "homeless", "homely", "honorable", "horrible", "hospitable", "hot", "huge", "hulking", "humdrum", "humorous", "hungry", "hurried", "hurt", "hushed", "husky", "hypnotic", "hysterical", "icky", "icy", "idiotic", "ignorant", "ill", "illegal", "illustrious", "imaginary", "immense", "imminent", "impartial", "imperfect", "important", "imported", "impossible", "incandescent", "incompetent", "inconclusive", "industrious", "inexpensive", "innate", "innocent", "inquisitive", "instinctive", "internal", "invincible", "irate", "itchy", "jaded", "jagged", "jazzy", "jealous", "jittery", "jobless", "jolly", "joyous", "judicious", "juicy", "jumbled", "jumpy", "juvenile", "kaput", "kind", "kindhearted", "knotty", "knowing", "knowledgeable", "known", "labored", "lackadaisical", "lacking", "lamentable", "languid", "large", "late", "laughable", "lavish", "lazy", "lean", "learned", "legal", "lethal", "level", "lewd", "light", "likeable", "literate", "little", "lively", "lonely", "long", "longing", "loose", "lopsided", "loud", "loutish", "lovely", "loving", "low", "lowly", "lucky", "ludicrous", "lush", "luxuriant", "lying", "lyrical", "macabre", "macho", "maddening", "madly", "magenta", "magical", "magnificent", "majestic", "makeshift", "malicious", "mammoth", "maniacal", "many", "marked", "massive", "materialistic", "mature", "measly", "meek", "melodic", "melted", "merciful", "mere", "mighty", "mindless", "miniature", "minor", "miscreant", "misty", "moaning", "modern", "moldy", "momentous", "motionless", "muddled", "muddy", "mundane", "murky", "mushy", "mute", "mysterious", "naive", "nappy", "narrow", "nasty", "naughty", "nauseating", "nebulous", "needless", "needy", "neighborly", "nervous", "new", "nice", "nifty", "noiseless", "noisy", "nonchalant", "nondescript", "nonstop", "nostalgic", "nosy", "noxious", "null", "numberless", "numerous", "nutritious", "nutty", "oafish", "obedient", "obeisant", "obnoxious", "obscene", "obsequious", "observant", "obsolete", "obtainable", "oceanic", "odd", "offbeat", "old", "omniscient", "onerous", "open", "optimal", "orange", "ordinary", "organic", "ossified", "outrageous", "outstanding", "oval", "overconfident", "overjoyed", "overrated", "overt", "overwrought", "painful", "painstaking", "panicky", "panoramic", "parched", "parsimonious", "pastoral", "pathetic", "peaceful", "penitent", "perfect", "periodic", "permissible", "perpetual", "petite", "phobic", "picayune", "piquant", "placid", "plain", "plastic", "plausible", "pleasant", "plucky", "pointless", "poised", "political", "poor", "possessive", "powerful", "precious", "premium", "pretty", "prickly", "productive", "profuse", "protective", "proud", "psychedelic", "psychotic", "puffy", "pumped", "puny", "purple", "purring", "puzzled", "quack", "quaint", "quarrelsome", "questionable", "quick", "quickest", "quiet", "quixotic", "quizzical", "rabid", "racial", "ragged", "rainy", "rambunctious", "rampant", "rapid", "rare", "raspy", "ratty", "real", "rebellious", "receptive", "recondite", "red", "redundant", "reflective", "relieved", "reminiscent", "repulsive", "resolute", "resonant", "rhetorical", "rich", "righteous", "rightful", "ripe", "ritzy", "roasted", "robust", "romantic", "roomy", "rotten", "rough", "round", "royal", "ruddy", "rural", "rustic", "ruthless", "sable", "sad", "salty", "sassy", "satisfying", "scandalous", "scarce", "scary", "scattered", "scientific", "scintillating", "scrawny", "screeching", "secretive", "sedate", "seemly", "selective", "selfish", "shaggy", "shaky", "shallow", "sharp", "shiny", "shivering", "shocking", "short", "shrill", "shy", "silent", "silky", "silly", "sincere", "skillful", "skinny", "sleepy", "slimy", "slippery", "sloppy", "slow", "small", "smelly", "smiling", "smoggy", "smooth", "sneaky", "snobbish", "snotty", "soft", "soggy", "solid", "somber", "sordid", "sore", "sour", "sparkling", "sparse", "spectacular", "spicy", "spiffy", "spiritual", "splendid", "spooky", "spotless", "spurious", "squalid", "square", "squealing", "squeamish", "staking", "stale", "standing", "statuesque", "steadfast", "steady", "steep", "stereotyped", "sticky", "stimulating", "stingy", "stormy", "straight", "strange", "strong", "stupid", "subdued", "subsequent", "substantial", "successful", "succinct", "sulky", "super", "supreme", "swanky", "sweet", "sweltering", "swift", "symptomatic", "synonymous", "taboo", "tacit", "tacky", "talented", "tall", "tame", "tan", "tangible", "tangy", "tart", "tasteful", "tasteless", "tasty", "tawdry", "tearful", "teeny", "telling", "temporary", "tender", "tense", "tenuous", "terrible", "tested", "testy", "thankful", "therapeutic", "thinkable", "thirsty", "thoughtful", "thoughtless", "threatening", "thundering", "tight", "tightfisted", "tiny", "tired", "tiresome", "toothsome", "torpid", "tough", "towering", "tranquil", "trashy", "tricky", "trite", "troubled", "truculent", "typical", "ubiquitous", "ugliest", "ugly", "ultra", "unable", "unaccountable", "unadvised", "unarmed", "unbecoming", "unbiased", "uncovered", "understood", "undesirable", "unequaled", "uneven", "uninterested", "unsightly", "unsuitable", "unusual", "upbeat", "uppity", "upset", "uptight", "used", "utopian", "utter", "uttermost", "vacuous", "vague", "various", "vast", "vengeful", "venomous", "verdant", "versed", "victorious", "vigorous", "vivacious", "voiceless", "volatile", "voracious", "vulgar", "wacky", "waggish", "wakeful", "wandering", "wanting", "warlike", "warm", "wary", "wasteful", "watchful", "watery", "weak", "wealthy", "weary", "wee", "wet", "whimsical", "whispering", "white", "wholesale", "wicked", "wide", "wild", "willing", "wiry", "wise", "wistful", "witty", "woebegone", "womanly", "wonderful", "wooden", "woozy", "workable", "worried", "worthless", "wrathful", "wretched", "wrong", "wry", "yellow", "yielding", "young", "youthful", "yummy", "zany", "zealous", "zippy"}
	adjCount = len(adjectives)

	nouns = []string{"able", "account", "achieve", "achiever", "acoustics", "act", "action", "activity", "actor", "addition", "adjustment", "advertisement", "advice", "aftermath", "afternoon", "afterthought", "agreement", "air", "airplane", "airport", "alarm", "alley", "amount", "amusement", "anger", "angle", "animal", "answer", "ant", "ants", "apparatus", "apparel", "apple", "apples", "appliance", "approval", "arch", "argument", "arithmetic", "arm", "army", "art", "attack", "attempt", "attention", "attraction", "aunt", "authority", "babies", "baby", "back", "badge", "bag", "bait", "balance", "ball", "balloon", "balls", "banana", "band", "base", "baseball", "basin", "basket", "basketball", "bat", "bath", "battle", "bead", "beam", "bean", "bear", "bears", "beast", "bed", "bedroom", "beds", "bee", "beef", "beetle", "beggar", "beginner", "behavior", "belief", "believe", "bell", "bells", "berry", "bike", "bikes", "bird", "birds", "birth", "birthday", "bit", "bite", "blade", "blood", "blow", "board", "boat", "boats", "body", "bomb", "bone", "book", "books", "boot", "border", "bottle", "boundary", "box", "boy", "boys", "brain", "brake", "branch", "brass", "bread", "breakfast", "breath", "brick", "bridge", "brother", "brothers", "brush", "bubble", "bucket", "building", "bulb", "bun", "burn", "burst", "bushes", "business", "butter", "button", "cabbage", "cable", "cactus", "cake", "cakes", "calculator", "calendar", "camera", "camp", "can", "cannon", "canvas", "cap", "caption", "car", "card", "care", "carpenter", "carriage", "cars", "cart", "cast", "cat", "cats", "cattle", "cause", "cave", "celery", "cellar", "cemetery", "cent", "chain", "chair", "chairs", "chalk", "chance", "change", "channel", "cheese", "cherries", "cherry", "chess", "chicken", "chickens", "children", "chin", "church", "circle", "clam", "class", "clock", "clocks", "cloth", "cloud", "clouds", "clover", "club", "coach", "coal", "coast", "coat", "cobweb", "coil", "collar", "color", "comb", "comfort", "committee", "company", "comparison", "competition", "condition", "connection", "control", "cook", "copper", "copy", "cord", "cork", "corn", "cough", "country", "cover", "cow", "cows", "crack", "cracker", "crate", "crayon", "cream", "creator", "creature", "credit", "crib", "crime", "crook", "crow", "crowd", "crown", "crush", "cry", "cub", "cup", "current", "curtain", "curve", "cushion", "dad", "daughter", "day", "death", "debt", "decision", "deer", "degree", "design", "desire", "desk", "destruction", "detail", "development", "digestion", "dime", "dinner", "dinosaurs", "direction", "dirt", "discovery", "discussion", "disease", "disgust", "distance", "distribution", "division", "dock", "doctor", "dog", "dogs", "doll", "dolls", "donkey", "door", "downtown", "drain", "drawer", "dress", "drink", "driving", "drop", "drug", "drum", "duck", "ducks", "dust", "ear", "earth", "earthquake", "edge", "education", "effect", "egg", "eggnog", "eggs", "elbow", "end", "engine", "error", "event", "example", "exchange", "existence", "expansion", "experience", "expert", "eye", "eyes", "face", "fact", "fairies", "fall", "family", "fan", "fang", "farm", "farmer", "father", "faucet", "fear", "feast", "feather", "feeling", "feet", "fiction", "field", "fifth", "fight", "finger", "fire", "fireman", "fish", "flag", "flame", "flavor", "flesh", "flight", "flock", "floor", "flower", "flowers", "fly", "fog", "fold", "food", "foot", "force", "fork", "form", "fowl", "frame", "friction", "friend", "friends", "frog", "frogs", "front", "fruit", "fuel", "furniture", "galley", "game", "garden", "gate", "geese", "ghost", "giants", "giraffe", "girl", "girls", "glass", "glove", "glue", "goat", "gold", "goldfish", "goodbye", "goose", "government", "governor", "grade", "grain", "grandfather", "grandmother", "grape", "grass", "grip", "ground", "group", "growth", "guide", "guitar", "gun", "hair", "haircut", "hall", "hammer", "hand", "hands", "harbor", "harmony", "hat", "hate", "head", "health", "hearing", "heart", "heat", "help", "hen", "hill", "history", "hobbies", "hole", "holiday", "home", "honey", "hook", "hope", "horn", "horse", "horses", "hose", "hospital", "hot", "hour", "house", "houses", "humor", "hydrant", "ice", "icicle", "idea", "impulse", "income", "increase", "industry", "ink", "insect", "instrument", "insurance", "interest", "invention", "iron", "island", "jail", "jam", "jar", "jeans", "jelly", "jellyfish", "jewel", "join", "joke", "journey", "judge", "juice", "jump", "kettle", "key", "kick", "kiss", "kite", "kitten", "kittens", "kitty", "knee", "knife", "knot", "knowledge", "laborer", "lace", "ladybug", "lake", "lamp", "land", "language", "laugh", "lawyer", "lead", "leaf", "learning", "leather", "leg", "legs", "letter", "letters", "lettuce", "level", "library", "lift", "light", "limit", "line", "linen", "lip", "liquid", "list", "lizards", "loaf", "lock", "locket", "look", "loss", "love", "low", "lumber", "lunch", "lunchroom", "machine", "magic", "maid", "mailbox", "man", "manager", "map", "marble", "mark", "market", "mask", "mass", "match", "meal", "measure", "meat", "meeting", "memory", "men", "metal", "mice", "middle", "milk", "mind", "mine", "minister", "mint", "minute", "mist", "mitten", "mom", "money", "monkey", "month", "moon", "morning", "mother", "motion", "mountain", "mouth", "move", "muscle", "music", "nail", "name", "nation", "neck", "need", "needle", "nerve", "nest", "net", "news", "night", "noise", "north", "nose", "note", "notebook", "number", "nut", "oatmeal", "observation", "ocean", "offer", "office", "oil", "operation", "opinion", "orange", "oranges", "order", "organization", "ornament", "oven", "owl", "owner", "page", "pail", "pain", "paint", "pan", "pancake", "paper", "parcel", "parent", "park", "part", "partner", "party", "passenger", "paste", "patch", "payment", "peace", "pear", "pen", "pencil", "person", "pest", "pet", "pets", "pickle", "picture", "pie", "pies", "pig", "pigs", "pin", "pipe", "pizzas", "place", "plane", "planes", "plant", "plantation", "plants", "plastic", "plate", "play", "playground", "pleasure", "plot", "plough", "pocket", "point", "poison", "police", "polish", "pollution", "popcorn", "porter", "position", "pot", "potato", "powder", "power", "price", "print", "prison", "process", "produce", "profit", "property", "prose", "protest", "pull", "pump", "punishment", "purpose", "push", "quarter", "quartz", "queen", "question", "quicksand", "quiet", "quill", "quilt", "quince", "quiver", "rabbit", "rabbits", "rail", "railway", "rain", "rainstorm", "rake", "range", "rat", "rate", "ray", "reaction", "reading", "reason", "receipt", "recess", "record", "regret", "relation", "religion", "representative", "request", "respect", "rest", "reward", "rhythm", "rice", "riddle", "rifle", "ring", "rings", "river", "road", "robin", "rock", "rod", "roll", "roof", "room", "root", "rose", "route", "rub", "rule", "run", "sack", "sail", "salt", "sand", "scale", "scarecrow", "scarf", "scene", "scent", "school", "science", "scissors", "screw", "sea", "seashore", "seat", "secretary", "seed", "selection", "self", "sense", "servant", "shade", "shake", "shame", "shape", "sheep", "sheet", "shelf", "ship", "shirt", "shock", "shoe", "shoes", "shop", "show", "side", "sidewalk", "sign", "silk", "silver", "sink", "sister", "sisters", "size", "skate", "skin", "skirt", "sky", "slave", "sleep", "sleet", "slip", "slope", "smash", "smell", "smile", "smoke", "snail", "snails", "snake", "snakes", "sneeze", "snow", "soap", "society", "sock", "soda", "sofa", "son", "song", "songs", "sort", "sound", "soup", "space", "spade", "spark", "spiders", "sponge", "spoon", "spot", "spring", "spy", "square", "squirrel", "stage", "stamp", "star", "start", "statement", "station", "steam", "steel", "stem", "step", "stew", "stick", "sticks", "stitch", "stocking", "stomach", "stone", "stop", "store", "story", "stove", "stranger", "straw", "stream", "street", "stretch", "string", "structure", "substance", "sugar", "suggestion", "suit", "summer", "sun", "support", "surprise", "sweater", "swim", "swing", "system", "table", "tail", "talk", "tank", "taste", "tax", "teaching", "team", "teeth", "temper", "tendency", "tent", "territory", "test", "texture", "theory", "thing", "things", "thought", "thread", "thrill", "throat", "throne", "thumb", "thunder", "ticket", "tiger", "time", "tin", "title", "toad", "toe", "toes", "tomatoes", "tongue", "tooth", "toothbrush", "toothpaste", "top", "touch", "town", "toy", "toys", "trade", "trail", "train", "trains", "tramp", "transport", "tray", "treatment", "tree", "trees", "trick", "trip", "trouble", "trousers", "truck", "trucks", "tub", "turkey", "turn", "twig", "twist", "umbrella", "uncle", "underwear", "unit", "use", "vacation", "value", "van", "vase", "vegetable", "veil", "vein", "verse", "vessel", "vest", "view", "visitor", "voice", "volcano", "volleyball", "voyage", "walk", "wall", "war", "wash", "waste", "watch", "water", "wave", "waves", "wax", "way", "wealth", "weather", "week", "weight", "wheel", "whip", "whistle", "wilderness", "wind", "window", "wine", "wing", "winter", "wire", "wish", "woman", "women", "wood", "wool", "word", "work", "worm", "wound", "wren", "wrench", "wrist", "writer", "writing", "yak", "yam", "yard", "yarn", "year", "yoke", "zebra", "zephyr", "zinc", "zipper"}
	nounCount = len(nouns)

	vowels = []string{"a", "e", "i", "o", "u"}

	rand.Seed(time.Now().Unix())
}

func isVowel(letter string) bool {
	for _, vowel := range vowels {
		if vowel == letter {
			return true
		}
	}
	return false
}

func toTitleCase(wordList []string) []string {
	var titleCaseWordList []string
	for _, word := range wordList {
		word = strings.ToLower(word)
		word = strings.Title(word)
		titleCaseWordList = append(titleCaseWordList, word)
	}
	return titleCaseWordList
}

// Adjective returns a single adjective.
// Returns a string.
func Adjective() string {
	return adjectives[rand.Intn(adjCount)]
}

// Noun returns a single noun.
// Returns a string.
func Noun() string {
	return nouns[rand.Intn(nounCount)]
}

// GenerateSpecial generates random strings of custom length, custom separator and optional title casing.
// Returns a string.
func GenerateSpecial(titleCase bool, wordCount int, separator string) string {
	if wordCount < 2 || wordCount > 10 {
		panic("wordCount must be between 2 and 10")
	}

	var phrase []string
	phrase = append(phrase, Adjective())
	phrase = append(phrase, Noun())

	if wordCount > 5 {
		for i := 0; i < wordCount-2; i++ {
			phrase = append([]string{Adjective()}, phrase...)
		}
	} else {

		if wordCount > 2 {
			phrase = append([]string{Adjective()}, phrase...)
		}

		if wordCount > 3 {
			if isVowel(string(phrase[0][0])) {
				article := []string{"an", "the"}[rand.Intn(2)]
				phrase = append([]string{article}, phrase...)
			} else {
				article := []string{"a", "the"}[rand.Intn(2)]
				phrase = append([]string{article}, phrase...)
			}
		}

		if wordCount > 4 {
			phrase = append(phrase, "")
			copy(phrase[3:], phrase[2:])
			phrase[2] = "and"
		}
	}

	if titleCase {
		phrase = toTitleCase(phrase)
	}

	return strings.Join(phrase, separator)
}

// Generate generates a readable random phrase of 2 words, each capitalized and no separator.
// Returns a string.
func Generate() string {
	return GenerateSpecial(true, 2, "")
}
