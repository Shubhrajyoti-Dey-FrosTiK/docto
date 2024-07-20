package constants

import (
	"time"
)

const DB = "portal"
const CONNECTION_STRING = "ATLAS_URI"
const REDIS_URI = "REDIS_URI"
const REDIS_USERNAME = "REDIS_USERNAME"
const REDIS_PASSWORD = "REDIS_PASSWORD"

const COLLECTION_STUDENT = "students"
const COLLECTION_GROUP = "groups"
const COLLECTION_CRITERIA = "criterias"
const COLLECTION_COMPANY = "companies"
const COLLECTION_OPPORTUNITY = "opportunities"
const COLLECTION_SLOT = "slots"
const COLLECTION_COMPANY_PROFILE = "companyProfiles"

const FIREBASE_PROJECT_ID = "FIREBASE_PROJECT_ID"

const CACHING_DURATION = 20 * time.Minute
const CACHE_CONTROL_HEADER = "Cache-Control"
const NO_CACHE = "no-cache"

// Keys of Cache
const GCP_JWKS = "GCP_JWKS"

const DB_PAGINATION = 30      // 30 results will be returned for DB pagination process
const DB_MAX_CYCLE_COUNT = 10 // The max number of cycles the DB will do to interupt the request

const TABLE_DOCTOR = "doctors"
const TABLE_PATIENT = "patients"

const ASSOCIATION_PATIENT = "Patients"
const ASSOCIATION_DOCTOR = "Doctors"
const ASSOCIATION_FILE = "Files"
