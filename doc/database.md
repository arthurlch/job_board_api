### DB Schema

erDiagram
    User ||--o{ JobSeeker : Owns
    User ||--o{ Company : Owns
    User ||--o{ Messages : Sender
    User ||--o{ Messages : Receiver
    User {
        int id
        string name
        string email
        string phone
        string role
        timestamp created_at
        timestamp updated_at
    }

    JobSeeker {
        int id
        int user_id
        string resume
        timestamp created_at
        timestamp updated_at
    }

    Company {
        int id
        int user_id
        string name
        string email
        string phone
        string website
        string logo
        text description
        timestamp created_at
        timestamp updated_at
    }

    Institution {
        int id
        string name
    }

    CompanyEntity {
        int id
        string name
    }

    JobCategory {
        int id
        string name
    }

    ExperienceType {
        int id
        string name
    }

    Education {
        int id
        int job_seeker_id
        int institution_id
        string degree
        string field_of_study
        date start_date
        date end_date
        timestamp created_at
        timestamp updated_at
    }

    Experience {
        int id
        int job_seeker_id
        string title
        int company_id
        string location
        date start_date
        date end_date
        int type_id
        text description
        timestamp created_at
        timestamp updated_at
    }

    Job {
        int id
        string title
        text description
        text requirements
        string location
        int salary
        int company_id
        int category_id
        job_status status
        timestamp created_at
        timestamp updated_at
    }

    Application {
        int id
        int job_seeker_id
        int job_id
        text cover_letter
        string resume
        application_status status
        timestamp created_at
        timestamp updated_at
    }

    Skill {
        int id
        string name
    }

    JobSeekerSkill {
        int job_seeker_id
        int skill_id
    }

    JobViews {
        int job_id
        int view_count
    }

    ChatbotInterview {
        int id
        int job_seeker_id
        int job_id
        interview_status status
        text review
        timestamp created_at
        timestamp updated_at
    }

    ChatbotConversation {
        int id
        int chatbot_interview_id
        string sender_type
        text content
        timestamp created_at
    }

    Messages {
        int id
        int sender_id
        int receiver_id
        text content
        string sender_type
        timestamp created_at
    }

    ScheduledInterview {
        int id
        int job_seeker_id
        int company_id
        timestamp scheduled_at
        string location
        text notes
        string meeting_link
        timestamp created_at
    }

    User }|..|{ JobSeeker : "Has"
    User }|..|{ Company : "Has"
    User }|..|{ Messages : "Sent"
    User }|..|{ Messages : "Received"
    JobSeeker }|--|| ExperienceType : "Has"
    JobSeeker }|--|| Education : "Has"
    JobSeeker }|..|{ Application : "Applied"
    JobSeeker }|..|{ JobSeekerSkill : "Has"
    JobSeeker }|--|| Experience : "Has"
    Experience }|--|| CompanyEntity : "At"
    Experience }|--|| ExperienceType : "Type"
    Company }|--|| CompanyEntity : "Has"
    Job }|--|| Company : "Posted by"
    Job }|--|| JobCategory : "Belongs to"
    Application }|--|| Job : "Application for"
    JobSeekerSkill }|--|| Skill : "Skill"
    Job }|..|{ JobViews : "Viewed"
    ChatbotInterview }|..|{ ChatbotConversation : "Has"
    JobSeeker }|..|{ ChatbotInterview : "Participated"
    Job }|..|{ ChatbotInterview : "Conducted for"
    User }|..|{ ScheduledInterview : "Participates"
    Company }|..|{ ScheduledInterview : "Schedules"
